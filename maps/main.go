package main

import (
	"fmt"
	"sync"
	"testGo/maps"
	"time"
)

const N = 100_000
const M = 50

var keys = [N]string{}

type foo struct {
	once sync.Once
}
func main() {
	var f foo = foo{}

	f.once.Do(func() {
		fmt.Println("Running once")
	})

	f.once.Do(func() {
		fmt.Println("Running twice")
	})

	//for i := 0 ; i < N ; i++ {
	//	key := fmt.Sprintf("%d", i)
	//	keys[i] = key
	//}
	//
	//// Set
	//fmt.Println("Set:")
	//inbuiltMapInsert()
	//customMapInsert()
	//customFastMapInsert()
	//
	//// Get
	//fmt.Println("\nGet:")
	//inbuiltMapGet()
	//customMapGet()
	//customFastMapGet()

	fmt.Println("Done!")
}

func inbuiltMapInsert(){

	start := time.Now()
	for i := 0; i < M; i++ {
		builtinMap := map[string]int{}
		for i := 0; i < N; i++ {
			builtinMap[keys[i]] = i
		}
	}

	fmt.Println("go map:", time.Now().Sub(start) / M)
}

func customMapInsert() {
	start := time.Now()
	for i := 0; i < M; i++ {
		stringIntMap := maps.NewStringIntOrderedMap()
		for i := 0; i < N; i++ {
			stringIntMap.Set(keys[i], i)
		}
	}

	fmt.Println("custom map:", time.Now().Sub(start) / M)
}

func customFastMapInsert() {
	start := time.Now()
	for i := 0; i < M; i++ {
		stringIntMap := maps.NewStringIntOrderedFastMap()
		for i := 0; i < N; i++ {
			stringIntMap.Set(keys[i], i)
		}
	}

	fmt.Println("fast map:", time.Now().Sub(start) / M)
}


func inbuiltMapGet(){
	builtinMap := map[string]int{}
	for i := 0; i < N; i++ {
		builtinMap[keys[i]] = i
	}

	start := time.Now()
	for i := 0; i < M; i++ {
		for i := 0; i < N; i++ {
			_ = builtinMap[keys[i]]
		}
	}

	fmt.Println("go map:", time.Now().Sub(start) / M)
}

func customMapGet() {
	stringIntMap := maps.NewStringIntOrderedMap()
	for i := 0; i < N; i++ {
		stringIntMap.Set(keys[i], i)
	}

	start := time.Now()
	for i := 0; i < M; i++ {
		for i := 0; i < N; i++ {
			stringIntMap.Get(keys[i])
		}
	}

	fmt.Println("custom map:", time.Now().Sub(start) / M)
}

func customFastMapGet() {
	stringIntMap := maps.NewStringIntOrderedFastMap()
	for i := 0; i < N; i++ {
		stringIntMap.Set(keys[i], i)
	}

	start := time.Now()
	for i := 0; i < M; i++ {
		for i := 0; i < N; i++ {
			stringIntMap.Get(keys[i])
		}
	}

	fmt.Println("fast map:", time.Now().Sub(start) / M)
}