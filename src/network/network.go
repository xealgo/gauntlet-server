package network

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
)

const (
	MaxInactivity   = 30 * time.Second
	CleanupInterval = 5 * time.Second
)

var (
	server     *net.UDPConn
	clients    = map[string]*Client{}
	clientLock sync.RWMutex
	sigs       = make(chan os.Signal, 1)
)

// Start begins listening for udp connections at a specified port.
func Start(port int, onConnect func(client *Client) error) error {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	service := fmt.Sprintf("127.0.0.1:%d", port)
	addr, err := net.ResolveUDPAddr("udp4", service)

	if err != nil {
		return err
	}

	// setup listener for incoming UDP connection
	if server, err = net.ListenUDP("udp", addr); err != nil {
		return err
	}

	log.Printf("server listening on port %d\n", port)

	defer server.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)

	go cleanup()
	go connect(onConnect)
	go shutdown(&wg)

	wg.Wait()
	return nil
}

// shutdown shuts the server down.
func shutdown(wg *sync.WaitGroup) {
	sig := <-sigs
	fmt.Println(sig)
	wg.Done()
}

// connect adds new clients to the client list.
func connect(onConnect func(client *Client) error) error {
	for {
		buffer := make([]byte, 1024)

		_, addr, err := server.ReadFromUDP(buffer)

		if err != nil {
			return err
		}

		if client, ok := clients[addr.String()]; ok {
			log.Printf("client %s says %s", client.UUID, string(buffer))
			client.LastPing = time.Now()
			Send(client, []byte("hello client!"))
			continue
		}

		client := Client{
			UUID:     uuid.New().String(),
			State:    ClientStateConnecting,
			Address:  addr,
			LastPing: time.Now(),
		}

		clientLock.Lock()
		clients[addr.String()] = &client
		clientLock.Unlock()

		log.Printf("client %s connected: %s, with message: %s\n", client.UUID, addr.String(), string(buffer))

		if err = onConnect(&client); err != nil {
			log.Printf("onConnect error: %s", err)
		}
	}
}

// cleanup this is very basic at the moment. A better approach would
// be to add each inactive client to a dead list which should typically
// be much lower than the active player count which will allow the lock
// to be released much sooner.
func cleanup() {
	log.Println("running cleanup cycle")

	// read locking will prevent both readers and writers
	clientLock.RLock()
	for key, client := range clients {
		delta := time.Now().Sub(client.LastPing)

		if delta > MaxInactivity {
			SendMessage(client, "You are being disconnected due to inactivity")
			clients[key] = nil
			log.Printf("removing client %s after %d seconds of inactivity\n", client.UUID, int(delta.Seconds()))
			delete(clients, key)
		}
	}
	clientLock.RUnlock()

	//fmt.Printf("THERE ARE %d ACTIVE PLAYERS\n", GetClientCount())
	time.Sleep(CleanupInterval)
	cleanup()
}

// GetClientCount TS method that returns the current client count.
func GetClientCount() uint {
	clientLock.RLock()
	defer clientLock.RUnlock()
	return uint(len(clients))
}

// Broadcast TS method that sends a message to all clients.
func Broadcast(packet []byte) error {
	clientLock.RLock()
	defer clientLock.RUnlock()

	for _, client := range clients {
		// we don't need to block because the routines will all continue
		// to run after the broadcast call has gone out of scope.
		go func(c *Client) {
			Send(c, packet)
		}(client)
	}

	return nil
}

// SendMessage wrapper for the Send method
func SendMessage(client *Client, message string) error {
	return Send(client, []byte(message))
}

// Send a packet to a particular client.
func Send(client *Client, packet []byte) error {
	if client != nil {
		if _, err := server.WriteToUDP(packet, client.Address); err != nil {
			return err
		}
	}
	// maybe return an error if the client IS nil?
	return nil
}
