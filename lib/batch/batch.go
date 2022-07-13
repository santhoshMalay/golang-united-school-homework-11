package batch

// package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	var userID int64 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	usersPerGoRoutine := int(n / pool)
	fmt.Printf("no.Of go routines %d \n", usersPerGoRoutine)
	//additionalUsers := int(n % pool)

	// wg.Add(int(pool))
	for i := 0; i < int(pool); i++ {
		wg.Add(1)
		go func(k int) {
			// fmt.Printf("Started GoRoutine %d \n", k)
			for j := k * usersPerGoRoutine; j < (k+1)*usersPerGoRoutine; j++ {
				tempUser := getOne(atomic.AddInt64(&userID, 1))
				mu.Lock()
				res = append(res, tempUser)
				mu.Unlock()
			}
			// fmt.Printf("Lenght of res is %d \n", len(res))
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}

func main() {
	x := getBatch(53, 5)
	fmt.Println(x)
	fmt.Printf("Lenght of x is %d", len(x))
}
