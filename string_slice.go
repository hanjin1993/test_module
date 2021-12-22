package test_module

import (
	"math/rand"
	"time"
)

func Random(maxNum int) int64 {
	//随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int64(r.Intn(maxNum))
}
