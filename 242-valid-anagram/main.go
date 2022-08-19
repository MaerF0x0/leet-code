func isAnagram(s string, t string) bool {
    sMap := make(map[rune]int)
    tMap := make(map[rune]int)
    
    wg := sync.WaitGroup{}
    wg.Add(2)
    var notAnagram bool
    go func() {
    for _, c := range []rune(s) { 
        sMap[c]++
    }
        wg.Done()
    }()
    
    go func() {
 for _, c := range []rune(t) { 
        tMap[c]++
    }
        wg.Done()
    }()
    wg.Wait()
 
    wg.Add(2)
    
    go func() {
    for k, count := range sMap {
        if count != tMap[k] {
            notAnagram = true
            break
        }
    }
        
        wg.Done()
    }()
    
    go func() {
    for k, count := range tMap {
        if count != sMap[k] {
          notAnagram = true
            break
        }
    } 
        
        wg.Done()
    }()
    wg.Wait()    
    
 return !notAnagram
}