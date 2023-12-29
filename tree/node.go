package tree

import (
	"encoding/json"
	"errors"
	"os"
)

type Node[T comparable] struct {
	Key    T          `json:"k"`
	Valide int        `json:"v"`
	Childs []*Node[T] `json:"c"`
}

func NewNode[T comparable](key T) *Node[T] {
	temp := new(Node[T])
	temp.Childs = []*Node[T]{}
	temp.Key = key
	return temp
}

func (this *Node[comparable]) HasChild(key comparable) (*Node[comparable], error) {

	for _, node := range this.Childs {
		if node.Key == key {
			return node, nil
		}
	}

	return this, errors.New("not found")
}

func (this *Node[comparable]) Merge(node *Node[comparable]) bool {
	if this.Key == node.Key {
		this.Childs = append(this.Childs, node.Childs...)
		return true
	}
	return false
}

func (this *Node[comparable]) Add(node *Node[comparable]) {
	for _, val := range this.Childs {
		if val.Merge(node) {
			return
		}
	}

	this.Childs = append(this.Childs, node)

}

func (this *Node[comparable]) Encode(path string) {

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

func (this *Node[string]) CheckWord(w []string) bool {

	next := this

	for _, l := range w {

		if temp, err := next.HasChild(l); err == nil {
			next = temp
		} else {
			return false
		}
	}

	return next.Valide == 1

}

func (this *Node[string]) CanCreateWord(w []string) bool {

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
