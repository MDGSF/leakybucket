package main

import (
	"fmt"
	"sync"

	".."
)

/*
Create one Store, it has ten map(safe map).
map's key is id, value is bucket, one id to one bucket.
*/

const Size = 10

// map <key: int, value: leakybucket.Bucket>
var Store []*sync.Map

func init() {
	Store = make([]*sync.Map, Size)
	for k := range Store {
		Store[k] = &sync.Map{}
	}
}

func CanPass(key int) bool {
	m := Store[key%Size]
	v, ok := m.Load(key)
	if !ok {
		m.Store(key, leakybucket.NewBucket(10, 100))
		v, _ = m.Load(key)
	}
	bucket := v.(*leakybucket.Bucket)
	return bucket.AddOne()
}

func main() {
	inputData := []int{1, 2019, 3000, 10000, 10, 78, 56, 1027}
	for i := 0; i < 20; i++ {
		inputData = append(inputData, 1)
	}

	for _, v := range inputData {
		ret := CanPass(v)
		fmt.Println(v, ret)
	}
}
