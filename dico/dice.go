package dico

type Dice []string

func (this Dice) Contains(letter string) bool {

	for _, l := range this {
		if l == letter {
			return true
		}
	}

	return false

}
