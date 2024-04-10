package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
// Write any import statements here
import "container/list"
func getMaximumEatenDishCount(N int32, D []int32, K int32) int32 {
  // This is like an LRU
// https://www.metacareers.com/profile/coding_puzzles?puzzle=958513514962507  
  
  lru := NewLRU(int(K))
 
  var platesEaten int32
  for _, d := range D {
    if lru.Has(d) {
      continue
    }
    lru.PushBack(d)
    platesEaten++
  }
 
  return platesEaten
}


type LRU struct {
  dataMap map[int32]*list.Element
  ll *list.List
  limit int
}

func NewLRU(limit int) *LRU {
  return &LRU{
    dataMap : make(map[int32]*list.Element, limit),
    ll : list.New(),
    limit: limit,
  }
}

func (lru *LRU) Has(i int32) bool {
  _, found := lru.dataMap[i]
  
  return found
}

// Push this dish to the back of the LRU
func (lru *LRU) PushBack(i int32) {
   if e, found := lru.dataMap[i] ; found {
      lru.ll.Remove(e)
   } else if lru.ll.Len() >= lru.limit {
      // Evict the oldest item
     removeElem := lru.ll.Front()
     delete(lru.dataMap, removeElem.Value.(int32))
     lru.ll.Remove(removeElem)
   }
  
  
    newElem := lru.ll.PushBack(i)
    lru.dataMap[i] = newElem   
}

