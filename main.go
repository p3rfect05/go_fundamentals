package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"

	"golang.org/x/exp/constraints"
)

func main() {
	a := []int{1, 2, 3, 13, 15}
	b := []int{4, 5, 6, 16, 14, 17}
	c := []int{7, 8, 9}
	for val := range mergeChannels(createChannel(a), createChannel(b), createChannel(c)) {
		fmt.Println(val)
	}
}

func task1() {
	fmt.Println("Hello world")
}

func task2() {
	a := 2
	c := true
	b := "lol"
	fmt.Printf("int: %v, bool: %v, string: %v", a, c, b)
}

func task3() {
	a := rand.Int63()
	if a%2 == 1 {
		fmt.Printf("%d is odd\n", a)
	} else {
		fmt.Printf("%d is even\n", a)
	}
}

func task4() {
	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i)
	}

}

type RealNumber interface {
	constraints.Integer | constraints.Float
}

func task5[T RealNumber](a, b T) T {
	return a + b
}

func task6[T RealNumber](sl []T) T {

	var sum T = 0
	for _, s := range sl {
		sum += s
	}
	return sum
}

func task7() {
	type Person struct {
		Name string
		Age  int
		City string
	}

	p := &Person{Name: "Me", Age: 10, City: "Moscow"}
	fmt.Println(p)
}

// task8 Swaps two number values
func task8[T any](a, b *T) {
	*a, *b = *b, *a
}

func task9[T RealNumber](a, b T) (T, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero (b == 0)")
	}
	return a / b, nil

}
func createChannel[T any](slice []T) chan T {
	ch := make(chan T)
	go func() {
		for _, val := range slice {
			ch <- val
		}
		close(ch)
	}()

	return ch
}

// task 10
func mergeChannels[T any](chans ...chan T) chan T {
	mergedChan := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	go func() {

		for _, ch := range chans {
			go func(ch chan T) {
				defer wg.Done()
				for val := range ch {
					mergedChan <- val
				}
			}(ch)
		}
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan
}
