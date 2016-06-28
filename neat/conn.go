package neat

import "math/rand"

type Conn struct {
	inNode   *Node   // input node
	outNode  *Node   // output node
	weight   float64 // connection weight
	disabled bool    // true if disabled
}

func NewConn(in, out *Node) *Conn {
	return &Conn{
		inNode:   in,
		outNode:  out,
		weight:   rand.Float64(),
		disabled: false,
	}
}

func (c *Conn) Output() {
	if !c.disabled {

	}
}
