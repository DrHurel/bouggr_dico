package tree

import (
	"encoding/json"
	"errors"
	"os"
)

type Node struct {
	Key    string  `json:"k"`
	Valide int     `json:"v"`
	Childs []*Node `json:"c"`
}

func (this *Node) HasChild(key string) (*Node, error) {

	for _, node := range this.Childs {
		if node.Key == key {
			return node, nil
		}
	}

	return this, errors.New("not found")
}

func (this *Node) Merge(node *Node) bool {
	if this.Key == node.Key {
		this.Childs = append(this.Childs, node.Childs...)
		return true
	}
	return false
}

func (this *Node) Add(node *Node) {
	for _, val := range this.Childs {
		if val.Merge(node) {
			return
		}
	}

	this.Childs = append(this.Childs, node)

}

func (this *Node) Encode(path string) {

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

func (this *Node) CheckWord(w string) bool {

	next := this

	for _, l := range w {

		if temp, err := next.HasChild(string(l)); err == nil {
			next = temp
		} else {
			return false
		}
	}

	return next.Valide == 1

}

func (this *Node) CanCreateWork(w string) bool {

	next := this

	for _, l := range w {

		if temp, err := next.HasChild(string(l)); err == nil {
			next = temp
		} else {
			return false
		}
	}

	return true

}
