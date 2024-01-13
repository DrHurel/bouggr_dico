package utils

func PrintGrid(Grid [4][4]rune) {
	for _, e := range Grid {
		for _, c := range e {
			print(string(c))
		}
		print("\n")
	}

}
