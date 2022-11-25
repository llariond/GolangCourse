package main

import (
	"fmt"
	"sync"
)

type safeMap struct {
	mut sync.Mutex
	m map[int64]int64 //make it template?
}

func NewSafeMap() (*safeMap, error) {
	return &safeMap{
		m : make(map[int64]int64),
	}, nil
}

func (sm *safeMap) At(key int64) (int64, bool) {
	sm.mut.Lock()
	defer sm.mut.Unlock()

	v, ok := sm.m[key]
	return v, ok
}

func (sm *safeMap) Add(key int64, val int64) {
	sm.mut.Lock()
	defer sm.mut.Unlock()

	sm.m[key] = val
}

func (sm *safeMap) Delete(key int64) {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	delete(sm.m, key)
}

func (sm *safeMap) GetKeys() []int64 {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	keys := make([]int64, len(sm.m))
	
	for k,_ := range(sm.m) {
		keys = append(keys, k)
	}

	return keys
}

func (sm *safeMap) GetValues() []int64 {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	vals := make([]int64, len(sm.m))
	
	for _,v := range(sm.m) {
		vals = append(vals, v)
	}

	return vals
}

func (sm *safeMap) GetRanges() ([]int64, []int64) {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	keys := make([]int64, len(sm.m))
	vals := make([]int64, len(sm.m))
	
	for k,v := range(sm.m) {
		keys = append(keys, k)
		vals = append(vals, v)
	}

	return keys, vals
}

func testRead(m *safeMap, n int64) {
	var cnt int64 = 0
    for {
        v,_ := m.At(cnt)
		v++ //just to make some use of variable
		cnt++
		if (cnt > n) {
			cnt = 0
		}
    }
}

func testWrite(m *safeMap, n int64) {
	var cnt int64 = 0
    for {
        m.Add(cnt, cnt)
		cnt++
		if (cnt > n) {
			cnt = 0
		}
    }
}

func main() {
	//just a placeholder
	//testing (start with -race)
	testMapP, _ := NewSafeMap() //a pointer to struct
}
