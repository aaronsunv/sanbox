package main
import (
 "fmt"
 "sync"
 "time"
)
const halfSecond = 500 * time.Millisecond
func main() {
 token := make(chan struct{}, 10)
 var wg sync.WaitGroup
 wg.Add(50)
for i := 0; i < 50; i++ {
  go func() {
   token <- struct{}{}
   defer wg.Done()
   time.Sleep(halfSecond)
   fmt.Printf("i is %d \n", i)
   <-token
  }()
 }
 start := time.Now()
 wg.Wait()
 fmt.Printf("It took me %d\n", time.Since(start))
}
