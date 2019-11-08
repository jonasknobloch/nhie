package category

import (
	"math/rand"
	"time"
)

func GetRandom() Category {
	categories := []Category{
		Harmless,
		Delicate,
		Offensive,
	}

	rand.Seed(time.Now().Unix())
	return categories[rand.Intn(len(categories))]
}
