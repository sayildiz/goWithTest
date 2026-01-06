package chsync

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// this says wait wantedCount times of wg.Done() signals before proceeding
		// WaitGroup is a counting semaphore used to wait for a group of goroutines to finish
		// Add(numberOfGoroutines) defines how many goRoutines to wait for and block until it reaches 0
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		// with wg.Go possible to do it without Add(wantedCount) and signaling Done
		// wg.Go will launch and finish the goroutines for you
		//for i := 0; i < wantedCount; i++ {
		//	wg.Go(func() {
		//		counter.Inc()
		//	})
		//}
		for i := 0; i < wantedCount; i++ {
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}

}
