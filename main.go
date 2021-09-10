package backoff

import (
   "math/rand"
   "time"

   "github.com/dblueman/nanolog"
)

const (
   min    = 3.
   max    = 15. * 60
   jitter = 0.1
)

type Backoff struct {
   nextDelay float64
}

func (b *Backoff) Delay() time.Duration {
   if b.nextDelay == 0. {
      b.Reset()
   }

   ret := b.nextDelay

   if b.nextDelay < max {
      b.nextDelay *= 3

      if b.nextDelay > max {
         b.nextDelay = max
      }
   }

   r := 2 * rand.Float64() - 1.
   s := ret * (1 + jitter * r)
   return time.Duration(s) * time.Second
}

func (b *Backoff) Sleep() {
   delay := b.Delay()
   nanolog.Info("retrying in %v", delay)
   time.Sleep(delay)
}

func (b *Backoff) Reset() {
   b.nextDelay = min
}
