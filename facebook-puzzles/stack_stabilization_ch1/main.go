func getMinimumDeflatedDiscCount(N int32, R []int32) int32 {
  var deflations int32

  var maxRadius int32 = R[N-1] 
  for i := len(R) -2 ; i >=0 ; i-- {
    // R[N-1]... R[i] is stable
    if R[i] >= maxRadius { 

       
      // we have to set R[i] to 1 less than maxRadius
      if maxRadius <= 1 {
        // Not allowed to set radius to 0
        return -1
      }
      maxRadius = maxRadius - 1
      deflations++
      continue
    }
    
    maxRadius = R[i] // This number is smaller than previously seen
  }
  
  
  return deflations
}
