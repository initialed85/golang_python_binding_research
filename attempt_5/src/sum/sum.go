package sum

import (
	"errors"
	"time"
	"sync"
	"log"
)

type SomeObject struct {
	Name   string
	Things []string
}

func NewSomeObject(name string) SomeObject {
	someObject := SomeObject{
		Name:   name,
		Things: []string{"hi", "how", "are", "you"},
	}

	return someObject
}

func Sum(a, b int) (float64, error) {
	if a == b {
		return 0.0, errors.New("oh bugger")
	}

	return float64(a) + float64(b), nil
}

func DoThing() int {
	things := []int{1,2,3,4,5,6,76,567,35476,5436,534}
	var locks = []sync.Mutex{}
	for range things {
		leChan := sync.Mutex{}
		leChan.Lock()
		locks = append(locks, leChan)
		go meDoingTheThing(leChan)
	}

	time.Sleep(5*time.Millisecond)

	for _, lock := range locks {
		lock.Lock()
		lock.Unlock()
	}

	return 1
}

func meDoingTheThing(blap sync.Mutex) {
	log.Println("Doing thing")
	<-time.After(5*time.Second)
	blap.Unlock()
	log.Println("No longer doing thing")
}