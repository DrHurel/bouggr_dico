package data_structure

import (
	"errors"
	"math"
)

type LangueCode int

func LangueCodeMap(lang []string) (res map[string]int) {

	res = make(map[string]int, len(lang))

	for i, l := range lang {
		res[l] = int(math.Pow(2, float64(i)))
	}

	return
}

func (l LangueCode) IsIn(code LangueCode) bool {

	return (code & l) > 0

}

type Node[T comparable, K any] struct {
	Key        T             `json:"k"`
	Value      K             `json:"v"`
	Children   []*Node[T, K] `json:"c"`
	Langue     LangueCode    `json:"l"`
	accessible bool
}

func NewNode[T comparable, K any](key T) *Node[T, K] {
	temp := new(Node[T, K])
	temp.Children = []*Node[T, K]{}
	temp.Key = key
	return temp
}

/*
Quick way to know if a node has children
*/
func (this *Node[T, K]) HasChildren() bool {
	return len(this.Children) != 0
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

	this.Children = append(this.Children, node)

}

// The `CheckWord` function is used to check if a given word exists in the tree and if its associated
// value matches a given test value.
func (this *Node[T, K]) CheckWord(w []T, test K, lang LangueCode, equal func(K, K) bool) bool {

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

	return equal(next.Value, test) && lang.IsIn(next.Langue)

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
