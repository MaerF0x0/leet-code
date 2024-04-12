package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"


func getMinCodeEntryTime(N int32, M int32, C []int32) int64 {

  var location int32 = 1
  var seconds int64  
  
  for i:= 0; i< len(C); i++ {
    a, b := pairs(location, C[i], N)
    fmt.Println(a, b)
    dist := int64(min(a, b))
    seconds += dist
//    fmt.Println("from location", location, "to ", C[i]," is", dist, "seconds->", seconds)
    
    location = C[i]
  }
  
  return seconds
}





func pairs(i, j, N int32) (int32, int32) {
  if i > j {
    return i-j , (N-(i-j))
  } else {
    return j-i, (N-(j-i))
  }
}



func min(i, j int32) int32 {
  if i < j {
    return i 
  } else {
    return j
  }
}
