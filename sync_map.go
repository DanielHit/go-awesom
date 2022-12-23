package service

import (
	"fmt"
	"sync"
)

func UseSyncMap() {
	a := &sync.Map{}
	a.Store("name", 2)
	a.Store("body", map[string]interface{}{
		"name": "yes",
		"age":  23,
	})

	fmt.Println(a.Load("name"))
	a.Range(func(key, value interface{}) bool {
		fmt.Println("key", key, "value", value)
		return true
	})
}

var a = DoWork
var b = DoWork

func GetRandInt(f func()) int64 {
	f()
	a("2", "3")
	return 1
}

func DoWork(a, b string) string {
	return a + b
}
