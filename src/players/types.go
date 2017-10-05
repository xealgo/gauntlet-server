package players

import "github.com/xealgo/gauntlet-server/src/network"

type Player struct {
	// position
	// velocity
	// etc.
	Name   string `json:"name"`
	Class  string `json:"class"`
	Skin   string `json:"skin"`
	Level  uint   `json:"level"`
	Health int    `json:"health"`
	Armor  int    `json:"armor"`

	// the client is only the network connection.
	client *network.Client
}
