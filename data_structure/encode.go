package data_structure

const (
	IS_A_WORD int32 = 1 << 8
)

type LangueCode int32

func LangueCodeMap(lang []string) (res map[string]int32) {

	res = make(map[string]int32, len(lang))

	for i, l := range lang {
		res[l] = 1 << (9 + i)
	}

	return
}

func Decode(code int32) rune {
	res := rune(0)
	for i := 0; i < 8; i++ {
		res += rune(code&(1<<i)) * (1 << i)
	}

	return res
}
