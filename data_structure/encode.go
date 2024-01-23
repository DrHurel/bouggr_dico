package data_structure

const (
	IS_A_WORD int16 = 1 << 8
)

type LangueCode int16

func LangueCodeMap(lang []string) (res map[string]int16) {

	res = make(map[string]int16, len(lang))

	for i, l := range lang {
		res[l] = 1 << (9 + i)
	}

	return
}

func Decode(code int16) rune {
	res := rune(0)
	for i := 0; i < 8; i++ {
		res += rune(code&(1<<i)) * (1 << i)
	}

	return res
}
