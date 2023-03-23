package utils

import (
	"math/rand"
	"time"
)

func RandShuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
