package qsort

import (
	"math/rand"
	"thelma/structs"
)

//QuickSort is function
func QuickSort(characters structs.Characters) structs.Characters {

	length := len(characters.List)

	if length <= 1 {
		return characters
	}

	m := characters.List[rand.Intn(length)]

	var less structs.Characters
	var middle structs.Characters
	var more structs.Characters

	for _, c := range characters.List {
		switch {
		case len(c.Name) > len(m.Name):
			less.List = append(less.List, c)
		case len(c.Name) == len(m.Name):
			middle.List = append(middle.List, c)
		case len(c.Name) < len(m.Name):
			more.List = append(more.List, c)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less.List = append(less.List, middle.List...)
	less.List = append(less.List, more.List...)

	return less
}
