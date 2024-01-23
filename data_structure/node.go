package data_structure

import (
	"errors"
)

type Node struct {
	_msgpack struct{} `msgpack:",as_array"`
	Value    int16    `json:"v"`
	Children []*Node  `json:"k"`
}

func NewNode(value int16) *Node {
	temp := new(Node)
	temp.Value = value
	return temp
}

func (this *Node) Convert() interface{} {
	a := []interface{}{
		this.Value,
	}

	if !this.HasChildren() {
		return this.Value
	}

	b := []interface{}{}

	for _, e := range this.Children {
		b = append(b, e.Convert())
	}

	a = append(a, b)

	return a

}

func (this *Node) AddLang(lang int16) {
	if this.Value&lang == 0 {
		this.Value += lang
	}
}

func (this *Node) SetAWord() {
	if this.Value&IS_A_WORD <= 0 {
		this.Value += IS_A_WORD
	}
}

/*
Quick way to know if a node has children
*/
func (this *Node) HasChildren() bool {
	return len(this.Children) != 0
}

func SameKey(val1, val2 int16) bool {
	//fmt.Printf("%s,%s,%s", string(Decode(val1)), string(Decode(val2)), "\n")
	for i := 0; i < 8; i++ {
		if val1&(1<<i) != val2&(1<<i) {
			return false
		}
	}
	return true
}

func (this *Node) GetChild(value rune) (*Node, error) {

	for _, node := range this.Children {

		if SameKey(node.Value, int16(value)) {
			return node, nil
		}
	}

	return this, errors.New("not found")
}

func (this *Node) Merge(node *Node) bool {
	if SameKey(this.Value, node.Value) {

		this.Children = append(this.Children, node.Children...)
		return true
	}
	return false
}

func (this *Node) Add(node *Node) {

	for _, val := range this.Children {
		if val.Merge(node) {
			return
		}
	}

	this.Children = append(this.Children, node)

}

// The `CheckWord` function is used to check if a given word exists in the tree and if its associated
// value matches a given test value.
func (this *Node) CheckWord(w string, lang int16) bool {

	next := this
	if len(w) < 3 {
		return false
	}

	for _, l := range w { // pour chaque lettre du mot

		if temp, err := next.GetChild(l); err == nil { // si on a un enfant qui correpond à la lettre
			next = temp //l'enfant est observé
		} else {
			return false
		}
	}
	//	afmt.Print(w, next.Value&IS_A_WORD > 0, "\n")
	return next.Value&IS_A_WORD > 0 && next.Value&lang > 0

}

func (this *Node) CanCreateWord(w string) bool {

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
