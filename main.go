package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex

func randSeq(n int) string {
	lock.Lock()
	defer lock.Unlock()
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func fillingMap() {
	m := make(map[string]string)
	for {
		m[randSeq(5)] = randSeq(10)
		fmt.Println(m)
		time.Sleep(time.Second)
	}
}

func outRandTen(p chan bool) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
	p <- true
}

func main() {
	p := make(chan bool)
	go outRandTen(p)
	go fillingMap()
	for {
		val := <-p
		if val {
			return
		}
		time.Sleep(time.Second)
	}
}
