package utils

import (
	"math/rand"
)

func Reduce[DT, RT any](slice []DT, fn func(el DT, ac RT) RT, initValue RT) RT {

	ac := initValue

	for _, v := range slice {
		ac = fn(v, ac)
	}
	return ac

}

func GetRandomCount(maxDuration int32) int32 {
	randCount := rand.Int31n(maxDuration + 1)
	if randCount == 0 {
		randCount = 1
	}
	return randCount
}
