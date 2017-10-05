package game

import "github.com/xealgo/gauntlet-server/src/sector"

var (
	sectors []sector.Sector
)

// Update updates each sector in a separate routine. This allows us to
// run multiple sectors in parallel on a single server.
func Update() {
	for _, sec := range sectors {
		go sector.Update(&sec)
	}
}
