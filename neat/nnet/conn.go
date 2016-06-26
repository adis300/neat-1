package nnet

import ()

type Conn struct {
	innum int          // innovation number
	from  *Node        // input node
	to    *Node        // output node
	ch    chan float64 // io channel
	w     float64      // connection weight
	dis   bool         // true if disabled
}

func NewConn(n int, in, out *Node) *Conn {
	return &Conn{
		innum: n,
		from:  in,
		to:    out,
		ch:    make(chan float64),
		w:     rand.Float64(),
		dis:   false,
	}
}

func (c *Conn) Output() {
}
