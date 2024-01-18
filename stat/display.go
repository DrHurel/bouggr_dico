package stat

import (
	"data_structure"
	"dico"
	"fmt"
	"math/rand"
	"time"
	"utils"
)

func Display(grid [4][4]rune,
	dices dico.Dices,
	r *rand.Rand,
	start time.Time,
	origin *data_structure.Node[rune, int],
) {
	var stat float64 = 0
	var distrition [17]int
	value := make([]float64, 0)
	min := 100
	max := 0
	var statSpeed time.Duration = 0

	for i := 0; i < 10000; i++ {

		grid = dices.Roll(r)
		start = time.Now()
		list := dico.AllWordInGrid(grid, origin)
		statSpeed += time.Since(start)
		n := len(list)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		stat += float64(n)
		value = append(value, float64(n))
		for _, e := range list {

			distrition[len(e)]++
		}
	}
	mean := float64(stat / 10000)

	fmt.Printf("min : %d | max %d | ecart-type %f | mean %f | avg speed %dms |Q1 %f Q2%f | Q3%f \n",
		min,
		max,
		mean,
		utils.EcartType(value, mean),
		statSpeed.Microseconds()/int64(stat),
		utils.NthTile(value, 1, 4),
		utils.NthTile(value, 2, 4),
		utils.NthTile(value, 3, 4),
	)
	for i, e := range distrition {
		fmt.Printf("Nombre de mots de longueur %d : %d \npourcentage  : %f%s\n", i, e, float64(e)*100/stat, "%")

	}

	fmt.Printf("Nombre total de mots %f", stat)
}