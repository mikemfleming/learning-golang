package kata

import (
  "strings"
  "fmt"
)

func duplicate_count(s1 string) int {
    var repeaters int
    counts := make(map[byte]int)
    
    s2 := strings.ToLower(s1)
    
    for i := 0; i < len(s2); i++ {
      counts[s2[i]] = 0
    }
    
    for j := 0; j < len(s2); j++ {
      counts[s2[j]]++
    }
    
    for _, v := range counts {
      if v > 1 {
        repeaters++
      }
    }
    
    fmt.Println(counts)
    
    return repeaters
  }



