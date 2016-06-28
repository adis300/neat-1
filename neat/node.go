package neat

type Node struct {
	incoming []*Conn // incoming connections
	outgoing []*Conn // outgoing connections
}

func NewNode(i int) *Node {
	return &Node{
		incoming: []*Conn{},
		outgoing: []*Conn{},
	}
}
