package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
	value int
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func(){
		instance = &singleton{value: 50}

	})
	return instance
}

func SingletonPattern() {
	s1 := GetInstance()
	fmt.Println("singleton pattern s1: ", s1.value)
	s2 := GetInstance()
	fmt.Println("singleton pattern s2: ", s2.value)

	if s1 == s2 {
		fmt.Println("singleton pattern is same")
	}
}

