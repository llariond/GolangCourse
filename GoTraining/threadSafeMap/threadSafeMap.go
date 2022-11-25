package main

import (
	"fmt"
	"sync"
)

type safeMap[K comparable, V any] struct {
	mut sync.Mutex
	m map[K]V
}

func NewSafeMap[K comparable, V any]() (*safeMap[K, V], error) {
	return &safeMap[K,V]{
		m : make(map[K]V),
	}, nil
}

func At[K comparable, V any](sm *safeMap[K,V], key K) (V, bool) {
	sm.mut.Lock()
	defer sm.mut.Unlock()

	v, ok := sm.m[key]
	return v, ok
}

func Add[K comparable, V any](sm *safeMap[K,V], key K, val V) {
	sm.mut.Lock()
	defer sm.mut.Unlock()

	sm.m[key] = val
}

func Delete[K comparable, V any](sm *safeMap[K,V], key K) {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	delete(sm.m, key)
}

func GetKeys[K comparable, V any](sm *safeMap[K,V]) []K {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	keys := make([]K, len(sm.m))
	
	for k,_ := range(sm.m) {
		keys = append(keys, k)
	}

	return keys
}

func GetValues[K comparable, V any](sm *safeMap[K,V]) []V {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	vals := make([]V, len(sm.m))
	
	for _,v := range(sm.m) {
		vals = append(vals, v)
	}

	return vals
}

func GetRanges[K comparable, V any] (sm *safeMap[K,V]) ([]K, []V) {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	keys := make([]K, len(sm.m))
	vals := make([]V, len(sm.m))
	
	for k,v := range(sm.m) {
		keys = append(keys, k)
		vals = append(vals, v)
	}

	return keys, vals
}

func main() {
	testMapP, _ := NewSafeMap[int64, int64]() //a pointer to struct

	Add[int64,int64](testMapP, 10, 1124)
	fmt.Println(At[int64, int64](testMapP, 10))
	fmt.Println(At[int64, int64](testMapP, 11))

	//QUESTION: HOW TO CHANGE FROM Add[type, type](....) to testMapP.Add() ??

	fmt.Print("END")
}
