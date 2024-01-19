package data_structure

import (
	"errors"
)

type Node[T comparable, K any] struct {
	Key        T             `json:"k"`
	Value      K             `json:"v"`
	Children   []*Node[T, K] `json:"c"`
	parent     *Node[T, K]
	accessible bool
}

func NewNode[T comparable, K any](key T) *Node[T, K] {
	temp := new(Node[T, K])
	temp.Children = []*Node[T, K]{}
	temp.Key = key
	return temp
}

func (this *Node[T, _]) Mark(w []T) {
	next := this

	for _, l := range w {
		next.accessible = true
		if temp, err := next.GetChild(l); err == nil {
			next = temp

		} else {
			return
		}
	}

}

/*
Quick way to know if a node has children
*/
func (this *Node[T, K]) HasChildren() bool {
	return len(this.Children) != 0
}

func (this *Node[T, K]) Parent() *Node[T, K] {
	return this.parent
}

func (this *Node[T, K]) GetChild(key T) (*Node[T, K], error) {

	for _, node := range this.Children {
		if node.Key == key {
			return node, nil
		}
	}

	return this, errors.New("not found")
}

func (this *Node[T, K]) Merge(node *Node[T, K]) bool {
	if this.Key == node.Key {
		for _, n := range node.Children {
			n.parent = this
		}
		this.Children = append(this.Children, node.Children...)
		return true
	}
	return false
}

func (this *Node[T, K]) Add(node *Node[T, K]) {

	for _, val := range this.Children {
		if val.Merge(node) {
			return
		}
	}
	node.parent = this
	this.Children = append(this.Children, node)

}

// The `CheckWord` function is used to check if a given word exists in the tree and if its associated
// value matches a given test value.
func (this *Node[T, K]) CheckWord(w []T, test K, equal func(K, K) bool) bool {

	next := this
	if len(w) < 3 {
		return false
	}

	for _, l := range w {

		if temp, err := next.GetChild(l); err == nil {
			next = temp
		} else {
			return false
		}
	}

	return equal(next.Value, test)

}

func (this *Node[T, _]) CanCreateWord(w []T) bool {

	next := this

	for _, l := range w {

		if temp, err := next.GetChild(l); err == nil {
			next = temp
		} else {
			return false
		}
	}

	return true

}
