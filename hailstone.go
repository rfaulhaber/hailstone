package main

import (
    "hailstone/node"
    "fmt"
    "os"
    "strconv"
)

func main() {
    arg := os.Args[1]
    n, err := strconv.Atoi(os.Args[1])
    nn := uint64(n)

    if err == nil {
        tree := hailstone(nn)
        fmt.Println("Depth: ", tree.TreeDepth())
    } else {
        fmt.Println("Invalid value", arg)
    }
}

func hailstone(n uint64) *node.Node {
    fmt.Println("Starting with: ", n)

    tree := node.MakeNode(n)
    n = nextHailstone(n)
    fmt.Println("Adding: ", n)
    tree.AddChild(n)
    treePtr := tree.Child

    for (n != 1) {
        treePtr.Child = node.MakeNode(nextHailstone(n))
        n = nextHailstone(n)
        fmt.Println("Adding: ", n)
        treePtr = treePtr.Child
    }

    return tree
}

func nextHailstone(n uint64) uint64 {
    if n % 2 == 0 {
        return n / 2
    } else {
        return (3 * n) + 1
    }
}
