package capture

import (
	"bes-chromie/src/encoder"
)

func (c *Capture) CreateEncoder() error {

	// defs
	var err error

	// start encoder
	c.Encoder, err = encoder.New()
	return err

}