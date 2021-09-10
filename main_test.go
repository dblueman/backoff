package backoff

import (
   "fmt"
   "testing"
)

func Test(t *testing.T) {
   var backoff Backoff

   for i := 0; i < 10; i++ {
      delay := backoff.Delay()
      fmt.Printf("delay %v\n", delay)
   }

   backoff.Reset()
   fmt.Println("reset")

   for i := 0; i < 10; i++ {
      delay := backoff.Delay()
      fmt.Printf("delay %v\n", delay)
   }
}
