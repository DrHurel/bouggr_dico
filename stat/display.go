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
	origin *data_structure.Node,
	noParsed *data_structure.Node,
	langs int32,
) {
	var stat float64 = 0
	stat2 := 0.
	var distrition [17]int
	var distrition2 [17]int
	value := make([]float64, 0)
	value2 := make([]float64, 0)
	min := 100
	max := 0
	min2 := 100
	max2 := 0
	var statSpeed time.Duration = 0
	nbError := 0
	missing := 0

	for i := 0; i < 10000; i++ {

		grid = dices.Roll(r)
		start = time.Now()
		list := dico.AllWordInGrid(grid, origin, langs)
		list2 := dico.AllWordInGrid(grid, noParsed, langs)
		if len(list) != len(list2) {
			missing += len(list2) - len(list)
			nbError++
		}
		statSpeed += time.Since(start)
		n := len(list)
		n2 := len(list2)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		if n2 < min2 {
			min2 = n2
		}
		if n2 > max2 {
			max2 = n2
		}
		stat += float64(n)
		stat2 += float64(n2)
		value = append(value, float64(n))
		value2 = append(value2, float64(n2))
		for _, e := range list {

			distrition[len(e)]++
		}
		for _, e := range list2 {

			distrition2[len(e)]++
		}
	}
	mean := float64(stat / 10000)
	mean2 := float64(stat2 / 10000)

	fmt.Printf("min : %d | max %d | ecart-type %f"+
		"| mean %f |avg speed %dms |Q1 %f Q2%f | Q3%f"+
		"|error count %d | missing %d | mean %f\n",
		min,
		max,
		mean,
		utils.EcartType(value, mean),
		statSpeed.Microseconds()/int64(stat),
		utils.NthTile(value, 1, 4),
		utils.NthTile(value, 2, 4),
		utils.NthTile(value, 3, 4),
		nbError,
		missing,
		float64(missing)/stat2,
	)

	fmt.Printf("min : %d | max %d | ecart-type %f"+
		"| mean %f |avg speed %dms |Q1 %f Q2%f | Q3%f"+
		"|error count %d | missing %d | mean %f\n",
		min2,
		max2,
		mean2,
		utils.EcartType(value2, mean2),
		statSpeed.Microseconds()/int64(stat2),
		utils.NthTile(value2, 1, 4),
		utils.NthTile(value2, 2, 4),
		utils.NthTile(value2, 3, 4),
		nbError,
		missing,
		float64(missing)/stat2,
	)
	for i, e := range distrition {
		fmt.Printf("Nombre de mots de longueur %d : %d "+
			"| %d \npourcentage  : %f%s\n", i, e, distrition2[i], float64(e)*100/stat, "%")

	}

	fmt.Printf("Nombre total de mots %f", stat)
}
