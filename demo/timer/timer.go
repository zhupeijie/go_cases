package timer

import (
	"fmt"
	"time"
)

type TickerX struct {
	ticker  *time.Ticker
	doFunc  func()
	doCount int
	gap     int // second
}

func NewTicker(gap int, dofun func()) *TickerX {
	return &TickerX{
		ticker:  time.NewTicker(time.Duration(gap) * time.Second),
		doFunc:  dofun,
		doCount: 0,
		gap:     gap,
	}
}

func (s TickerX) Tick() {
	for _ = range s.ticker.C {
		s.doFunc()
		s.doCount++
		if s.doCount == 10 {
			s.ticker.Reset(time.Duration(s.gap+1) * time.Second)
			s.gap++
			s.doCount = 0
			fmt.Println("-----------here is reset ticker ", s.gap)
		}
	}
}

func (s TickerX) Stop() {
	fmt.Println("here is stop ticker")
	s.ticker.Stop()
}
