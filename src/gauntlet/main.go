package main

import (
	"fmt"

	"github.com/xealgo/gauntlet-server/src/network"
)

func main() {
	// setup the game world.
	// @TODO: Get the offset and max from args
	// game.Preload(0, 10)
	//
	// // now begin the main game logic loop.
	// game.Update()

	// finally, start listening for incoming connections.
	network.Start(1386, func(client *network.Client) error {
		count := network.GetClientCount()
		network.Send(client, []byte(fmt.Sprintf("welcome! You are player %d\n", count)))
		network.Broadcast([]byte(fmt.Sprintf("player %d has joined the server\n", count)))
		return nil
	})
}
