package sector

import "github.com/xealgo/gauntlet-server/src/players"

// Sector represents a particular map within the game world.
type Sector struct {
	//sync.RWMutex

	UID  string `json:"unique-id"`
	Name string `json:"name"`
	Cell uint   `json:"cell"` // the virtual grid node this sector has been assigned to.

	Players []*players.Player
}
