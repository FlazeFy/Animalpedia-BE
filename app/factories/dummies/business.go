package dummies

import (
	"math/rand"
	"time"
)

func DummyDctType() string {
	seed := []string{
		"animal_zone",
		"animal_status",
		"animal_category",
		"animal_region",
	}

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(seed))
	res := seed[idx]

	return res
}
