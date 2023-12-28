package dico

type Dice []rune

func (this Dice) Contains(letter rune) bool {

	for _, l := range this {
		if l == letter {
			return true
		}
	}

	return false

}
