package golib

import "fmt"

type Tree[T any] struct {
	Core        T
	Children    []*Tree[T]
	parent      *Tree[T]
	parentIndex int
}

var depthMap [1024]bool

func (t Tree[T]) String() string {
	var s = ""
	t.stringNode(0, &s)
	return s
}

func (node *Tree[T]) stringNode(depth int, buffer *string) {
	depthMap[depth] = true
	for c := 0; c < depth; c++ {
		if depthMap[c] {
			if c == depth-1 {
				*buffer += "\033[1;37m|--\033[m"
			} else {
				*buffer += "\033[1;37m|  \033[m"
			}
		} else {
			*buffer += "\033[1;37m   \033[m"
		}
	}
	*buffer += fmt.Sprintf("\033[0;35m%v\033[m\n", node.Core)

	for _, child := range node.Children {
		child.stringNode(depth+1, buffer)
	}
}
