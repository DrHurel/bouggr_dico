package tree

import (
	"encoding/json"
	"errors"
	"os"
)

type Node[T comparable, K any] struct {
	Key    T             `json:"k"`
	Value  K             `json:"v"`
	Childs []*Node[T, K] `json:"c"`
}

func NewNode[T comparable, K any](key T) *Node[T, K] {
	temp := new(Node[T, K])
	temp.Childs = []*Node[T, K]{}
	temp.Key = key
	return temp
}

func (this *Node[T, K]) HasChild(key T) (*Node[T, K], error) {

	for _, node := range this.Childs {
		if node.Key == key {
			return node, nil
		}
	}

	return this, errors.New("not found")
}

func (this *Node[T, K]) Merge(node *Node[T, K]) bool {
	if this.Key == node.Key {
		this.Childs = append(this.Childs, node.Childs...)
		return true
	}
	return false
}

func (this *Node[T, K]) Add(node *Node[T, K]) {
	for _, val := range this.Childs {
		if val.Merge(node) {
			return
		}
	}

	this.Childs = append(this.Childs, node)

}

func (this *Node[_, _]) Encode(path string) {

	if file, err := os.Create(path); err != nil {
		panic(err)
	} else {
		encoder := json.NewEncoder(file)
		err := encoder.Encode(this)
		if err != nil {
			panic(err)
		}
	}

}

func (this *Node[T, K]) CheckWord(w []T, test K, equal func(K, K) bool) bool {

	next := this

	for _, l := range w {

		if temp, err := next.HasChild(l); err == nil {
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

		if temp, err := next.HasChild(l); err == nil {
			next = temp
		} else {
			return false
		}
	}

	return true

}
