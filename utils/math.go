package utils

import (
	"math/rand"
	"time"
)

var (
	ran *rand.Rand
)

//RandomFloat64 return a random float64 num based on current time
func RandomFloat64() float64 {
	return ran.Float64()
}

//RandomFloat32 return a random float32 num based on current time
func RandomFloat32() float32 {
	return ran.Float32()
}

//RandomIntn return a random int based on current time with a given maxnum
func RandomIntn(max int) int {
	return ran.Intn(max)
}

func RandomInt63n(max int64) int64 {
	return ran.Int63n(max)
}

func init() {
	// 2^9 = 512
	seed := time.Now().UnixNano() >> 9
	ran = rand.New(rand.NewSource(seed))
}
