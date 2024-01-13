package data_structure

type Pile[T any] struct {
	value []T
	len   uint
}

func (pile Pile[T]) Empiler(e T) {
	pile.value = append(pile.value, e)
	pile.len++
}

func (pile Pile[T]) Depiler() T {
	res := pile.value[0]
	pile.value = pile.value[1:]
	pile.len--
	return res
}

func (pile Pile[T]) Len() uint {
	return pile.len
}
