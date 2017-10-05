package network

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type NetworkTests struct{}

var (
	_ = Suite(&NetworkTests{})
)

func (t *NetworkTests) TestNewSector(c *C) {

}
