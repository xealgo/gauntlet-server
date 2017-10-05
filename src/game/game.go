package game

import "github.com/xealgo/gauntlet-server/src/sector"

var (
	sectors []sector.Sector
)

// Preload loads a subset of the current sector map files starting at offset until
// max. This will allow us to scale up and tweak how many servers are needed to run
// the entire game world. The actual deployment / scaling stret will be taken care of
// externally to the server - this will allow us to create multiple "realms" as well.
func Preload(offset, max int) {
	// fetch sector list....
	// storage.Get(SectorListFile)

	// now, load each sectors map data..
	for _, sec := range sectors {
		sector.Load(&sec)
	}
}

// Update updates each sector in a separate routine. This allows us to
// run multiple sectors in parallel on a single server.
func Update() {
	for _, sec := range sectors {
		go sector.Update(&sec)
	}
}
