package node

type Node struct {
    Value uint64
    Child *Node
}

func MakeNode(v uint64) *Node {
    return &Node{v, nil}
}

func (n *Node) AddChild(v uint64) {
    n.Child = MakeNode(v)
}

func (n *Node) TreeContains(v uint64) (bool, *Node) {
    if n.Value == v {
        return true, n
    } else if n.Child == nil {
        return false, nil
    } else {
        return n.Child.TreeContains(v)
    }
}

func (n *Node) TreeDepth() int {
    if n.Child != nil {
        return 1 + n.Child.TreeDepth()
    } else {
        return 0
    }
}
