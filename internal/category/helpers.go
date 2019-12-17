package category

import (
	"math/rand"
)

func GetRandom(categories ...Category) Category {
	if len(categories) == 0 {
		categories = []Category{
			Harmless,
			Delicate,
			Offensive,
		}
	}

	length := len(categories)

	if length == 1 {
		return categories[0]
	}

	return categories[rand.Intn(length)]
}
