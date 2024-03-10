package exchange

import (
	"context"
	container "exchange/exchange/pkg/OHLCV_repo"
	"exchange/exchange/pkg/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	dummyData = []service.OHLCV{
		{
			Open:   1,
			High:   10,
			Low:    3,
			Close:  5,
			Volume: 100,
		},
		{
			Open:   2,
			High:   20,
			Low:    2,
			Close:  7,
			Volume: 100,
		},
	}
)

type Exchange struct {
	Stats *container.Container
}

func (e *Exchange) Statistic(broker *service.BrokerID, stream service.Exchange_StatisticServer) error {
	ch := make(chan *service.OHLCV)
	e.Stats.AddChan(ch)
	for {
		select {
		case stat := <-ch:
			if stat != nil {
				stream.Send(stat)
			} else {
				e.Stats.DeleteChan(ch)
				return grpc.Errorf(codes.DataLoss, "")
			}
		case <-e.Stats.FinishCtx.Done():
			e.Stats.DeleteChan(ch)
			return grpc.Errorf(codes.OK, "")
		case <-stream.Context().Done():
			e.Stats.DeleteChan(ch)
			return grpc.Errorf(codes.OK, "")
		}
	}
	e.Stats.DeleteChan(ch)
	return grpc.Errorf(codes.OK, "")
}

func (e *Exchange) Create(ctx context.Context, deal *service.Deal) (*service.DealID, error) {
	id := &service.DealID{
		ID:                   deal.ID,
		BrokerID:             int64(deal.BrokerID),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	err := e.Stats.AddDeal(id, deal)
	if err != nil {
		return nil,grpc.Errorf(codes.AlreadyExists, "")
	}
	return id, grpc.Errorf(codes.OK, "")
}

func (e *Exchange) Cancel(ctx context.Context, dealId *service.DealID) (*service.CancelResult, error) {
	id := &service.DealID{
		ID:                   dealId.ID,
		BrokerID:             int64(dealId.BrokerID),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := &service.CancelResult{
		Success:              true,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	err := e.Stats.CancelDeal(id)
	if err != nil {
		res.Success = false
	}
	return res, grpc.Errorf(codes.OK, "")
}

func (e *Exchange) Results(broker *service.BrokerID, stream service.Exchange_ResultsServer) error {
	ch := make(chan *service.Deal)
	e.Stats.AddChanRes(ch)
	for {
		select {
		case res := <-ch:
			if res != nil {
				stream.Send(res)
			} else {
				e.Stats.DeleteChanRes(ch)
				return grpc.Errorf(codes.DataLoss, "")
			}
		case <-e.Stats.FinishCtx.Done():
			e.Stats.DeleteChanRes(ch)
			return grpc.Errorf(codes.OK, "")
		case <-stream.Context().Done():
			e.Stats.DeleteChanRes(ch)
			return grpc.Errorf(codes.OK, "")
		}
	}
	e.Stats.DeleteChanRes(ch)
	return nil
}

func StartExchangeServer(ctx context.Context, listenAddr string) error {
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	server := grpc.NewServer()
	exchange := &Exchange{Stats: container.CreateStatsContainer(dummyData)}
	service.RegisterExchangeServer(server, exchange)
	go func(lisAdr net.Listener) {
		if err := server.Serve(lisAdr); err != nil {
			fmt.Println("error to start server!")
		}
	}(lis)
	go func() {
		for {
			select {
			case <-ctx.Done():
				exchange.Stats.ClearStatsContainer()
				server.GracefulStop()
				return
			}
		}
	}()
	return nil
}
