package exchange

import (
	context "context"
	"exchange/exchange/pkg/service"
	"fmt"
	"io"
	"log"
	"sync"
	"testing"
	"time"

	grpc "google.golang.org/grpc"
)

const listenAddr string = "127.0.0.1:8082"

func wait(amout int) {
	time.Sleep(time.Duration(amout) * 10 * time.Millisecond)
}

func getGrpcConn(t *testing.T) *grpc.ClientConn {
	grcpConn, err := grpc.Dial(
		listenAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("cant connect to grpc: %v", err)
	}
	return grcpConn
}

func TestStat(t *testing.T) {
	ctx, finish := context.WithCancel(context.Background())
	err := StartExchangeServer(ctx, listenAddr)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	wait(1)
	defer func() {
		finish()
		wait(2)
	}()
	conn := getGrpcConn(t)
	defer conn.Close()
	exchClient := service.NewExchangeClient(conn)
	if err != nil {
		fmt.Println(err)
	}
	wait(1)
	statStream1, err := exchClient.Statistic(context.Background(), &service.BrokerID{ID: 1})
	wait(1)
	statStream2, err := exchClient.Statistic(context.Background(), &service.BrokerID{ID: 2})

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			stat, err := statStream1.Recv()
			if err != nil && err != io.EOF {
				fmt.Printf("unexpected error %v\n", err)
				return
			} else if err == io.EOF || stat == nil {
				t.Errorf("invalid data")
			}
			log.Println("stat1", stat.String(), err)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			stat, err := statStream2.Recv()
			if err != nil && err != io.EOF {
				fmt.Printf("unexpected error %v\n", err)
				return
			} else if err == io.EOF || stat == nil {
				t.Errorf("invalid data")
			}
			log.Println("stat2", stat.String(), err)
		}
	}()
	wg.Wait()
}

func TestResults(t *testing.T) {
	ctx, finish := context.WithCancel(context.Background())
	err := StartExchangeServer(ctx, listenAddr)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	wait(1)
	defer func() {
		finish()
		wait(2)
	}()
	conn := getGrpcConn(t)
	defer conn.Close()
	exchClient := service.NewExchangeClient(conn)
	if err != nil {
		fmt.Println(err)
	}
	wait(1)
	resStream, err := exchClient.Results(context.Background(), &service.BrokerID{ID: 1})
	if err != nil {
		fmt.Println(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++  {
			res, err := resStream.Recv()
			if err != nil && err != io.EOF {
				fmt.Printf("unexpected error %v\n", err)
				return
			} else if err == io.EOF || res == nil {
				fmt.Println("invalid data")
			}
			log.Println("Res Deal", res.String())
		}
	}()
	dealToClose := &service.Deal{
		ID:                   1,
		BrokerID:             1,
		ClientID:             1,
		Ticker:               "",
		Amount:               100,
		Partial:              false,
		Time:                 1,
		Price:                5,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	exchClient.Create(context.Background(), dealToClose)
	exchClient.Create(context.Background(), &service.Deal{ID: 2, Ticker: "Test2"})
	exchClient.Create(context.Background(), &service.Deal{ID: 3, Ticker: "Test3"})
	wg.Wait()
}

func TestPartial(t *testing.T) {
	ctx, finish := context.WithCancel(context.Background())
	err := StartExchangeServer(ctx, listenAddr)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	defer func() {
		finish()
		wait(2)
	}()
	conn := getGrpcConn(t)
	defer conn.Close()
	exchClient := service.NewExchangeClient(conn)

	resStream, err := exchClient.Results(context.Background(), &service.BrokerID{ID: 1})
	if err != nil {
		t.Errorf("unexpected error")
	}
	statStream, err := exchClient.Statistic(context.Background(), &service.BrokerID{ID: 1})
	if err != nil {
		t.Errorf("unexpected error")
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		stat, err := statStream.Recv()
		exchClient.Create(context.Background(), &service.Deal{BrokerID: 1, Amount: int32(stat.Volume) + 1})
		res, err := resStream.Recv()
		if res.Partial != true {
			t.Errorf("expected partial flag")
		}
		if err != nil && err != io.EOF {
			fmt.Printf("unexpected error %v\n", err)
			return
		} else if err == io.EOF {
			fmt.Println("eof)")
		} else if res == nil {
			t.Errorf("invalid data recieved")
		}
		log.Println("Res Deal", res.String())
	}()
	wg.Wait()
}

func TestCancel(t *testing.T) {
	ctx, finish := context.WithCancel(context.Background())
	err := StartExchangeServer(ctx, listenAddr)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	defer func() {
		finish()
		wait(2)
	}()
	conn := getGrpcConn(t)
	defer conn.Close()
	exchClient := service.NewExchangeClient(conn)

	resStream, err := exchClient.Results(context.Background(), &service.BrokerID{ID: 1})
	if err != nil {
		t.Errorf("unexpected error")
	}
	statStream, err := exchClient.Statistic(context.Background(), &service.BrokerID{ID: 1})
	if err != nil {
		t.Errorf("unexpected error")
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wait(10)
	go func() {
		defer wg.Done()
		for {
			stat, err := statStream.Recv()
			if err != nil {
				t.Error("unexpected error create:", err)
			}
			dealId, err := exchClient.Create(context.Background(), &service.Deal{ID: 1, BrokerID: 1, Price: stat.Close - 1})
			if err != nil {
				t.Error("unexpected error create:", err)
			}
			cancelRes, err := exchClient.Cancel(context.Background(), dealId)
			log.Println(cancelRes.String())
			if err != nil {
				t.Error("unexpected error create:", err)
			}
			res, err := resStream.Recv()
			if err != nil && err != io.EOF {
				fmt.Printf("unexpected error %v\n", err)
				return
			} else if err == io.EOF {
				break
			}
			log.Println("Res Deal", res.String())
		}

	}()
	wg.Wait()
}
