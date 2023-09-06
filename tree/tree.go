package tree

type Tree[T any] struct {
	Children []Tree[T]
	Core     T
}
