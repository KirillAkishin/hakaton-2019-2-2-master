package container

import (
	"context"
	"exchange/exchange/pkg/service"
	"fmt"
	"sync"
	"time"
)

type Container struct {
	data 		map[string][]service.OHLCV
	mu           *sync.Mutex
	FinishCtx    context.Context
	finishFunc   context.CancelFunc
	currentTimeInSecs int
	currentStats *service.OHLCV
	statsChans   map[chan *service.OHLCV]struct{}
	resultsChans map[chan *service.Deal]struct{}
	pendingDeals map[string]*service.Deal
}

func CreateStatsContainer(rts []service.OHLCV) *Container {
	ctxFinish, finish := context.WithCancel(context.Background())
	c := Container{
		data:         make(map[string][]service.OHLCV, 0),
		mu:           &sync.Mutex{},
		FinishCtx:    ctxFinish,
		finishFunc:   finish,
		currentTimeInSecs: 0,
		currentStats: nil,
		resultsChans: make(map[chan *service.Deal]struct{}),
		statsChans:   make(map[chan *service.OHLCV]struct{}, 0),
		pendingDeals: make(map[string]*service.Deal, 0),
	}
	c.AddData("SPFB.RTS", rts)
	go func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		for {
			select {
			case <-ctx.Done():
				biddingIsAvailable := c.updateStats("SPFB.RTS")
				if !biddingIsAvailable {
					ctxFinish.Done()
					return
				}
				ctx, _ = context.WithTimeout(context.Background(), time.Second)
			case <-ctxFinish.Done():
				return
			}
		}
	}()
	return &c
}

func (c *Container) ClearStatsContainer() {
	for ch, _ := range c.statsChans {
		close(ch)
	}
	for ch, _ := range c.resultsChans {
		close(ch)
	}
	c.finishFunc()
}

func (c *Container) AddData(name string, data []service.OHLCV) {
	c.data[name] = data
}

func (c *Container) checkDeal(dealId string) (bool, bool) {
	deal := c.pendingDeals[dealId]
	if c.currentStats.Close <= deal.Price {
		if c.currentStats.Volume >= uint32(deal.Amount) {
			c.currentStats.Volume -= uint32(deal.Amount)
			return true, false
		} else {
			c.currentStats.Volume = 0
			deal.Amount -= int32(c.currentStats.Volume)
			deal.Partial = true
			return true, true
		}
	}
	return false, false
}

func (c *Container) updateStats(ticker string) bool {
	c.mu.Lock()
	defer func() {c.currentTimeInSecs++}()
	defer c.mu.Unlock()
	if c.currentTimeInSecs == len(c.data[ticker]){
		return false
	}
	c.currentStats = &c.data[ticker][c.currentTimeInSecs]
	for ch, _ := range c.statsChans {
		ch <- c.currentStats
	}
	toDelete := make([]string, 0)
	for key, val := range c.pendingDeals {
		matched, isPartial := c.checkDeal(key)
		if  matched {
			for ch, _ := range c.resultsChans {
				ch <- val
			}
		}
		if isPartial {
			toDelete = append(toDelete, key)
		}
	}
	for _, key := range toDelete {
		delete(c.pendingDeals, key)
	}
	return true
}

func (c *Container) AddChan(ch chan *service.OHLCV) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.statsChans[ch] = struct{}{}
}

func (c *Container) AddChanRes(ch chan *service.Deal) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.resultsChans[ch] = struct{}{}
}

func (c *Container) DeleteChan(ch chan *service.OHLCV) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.statsChans, ch)
}

func (c *Container) DeleteChanRes(ch chan *service.Deal) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.resultsChans, ch)
}

func getKey(deal *service.DealID) string {
	return fmt.Sprintf("%d_%d", deal.BrokerID, deal.ID)
}

func (c *Container) AddDeal(dealId *service.DealID, deal *service.Deal) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	key := getKey(dealId)
	if _, ok := c.pendingDeals[key]; !ok {
		c.pendingDeals[key] = deal
		return nil
	}
	return fmt.Errorf("deal in container")
}

func (c *Container) CancelDeal(deal *service.DealID) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	key := getKey(deal)
	if _, ok := c.pendingDeals[key]; ok {
		delete(c.pendingDeals, key)
		return nil
	}
	return fmt.Errorf("deal isn't in container")
}
