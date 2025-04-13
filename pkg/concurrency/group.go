package concurrency

import "sync"

type AsyncGroup struct {
	wg sync.WaitGroup
}

func (g *AsyncGroup) Do(fn func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		fn()
	}()
}

func (g *AsyncGroup) Wait() {
	g.wg.Wait()
}
