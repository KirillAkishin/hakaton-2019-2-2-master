package main

import (
	"broker/pkg/clients"
	"broker/pkg/handlers"
	"broker/pkg/middleware"
	"broker/pkg/order_history"
	"broker/pkg/requests"
	"broker/pkg/service"
	"broker/pkg/session"
	"broker/pkg/stats"
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println()
	rand.Seed(time.Now().UnixNano())
	templates := template.Must(template.ParseFiles("./template/index.html"))
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync() // flushes buffer, if any
	logger := zapLogger.Sugar()
	staticHandler := http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./template/static/")),
	)
	http.Handle("/static/", staticHandler)

	dsn := "root:qwerty@tcp(localhost:3306)/hak?"
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"
	db, err := sql.Open("mysql", dsn)
	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	grcpConn, err := grpc.Dial(
		listenAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("cant connect to grpc: %v", err)
	}
	statRepo := stats.NewRepository(db)
	exchClient := service.NewExchangeClient(grcpConn)
	statStream1, err := exchClient.Statistic(context.Background(), &service.BrokerID{ID: 1})
	
	go func() {
		lastSec := -1
		var lastPrice float32
		var secondMax float32
		var secondMin float32
		var curPrice float32
		var secondVolume int
		var volume int
		var allVolume uint32
		ch := chan *service.OHLCV
		var stat *service.OHLCV
		go func(chanel chan *service.OHLCV) {
			stat, err := statStream1.Recv()
			if err != nil && err != io.EOF {
				fmt.Printf("unexpected error %v\n", err)
				return
			}
			chanel <- stat
		}(ch)
		ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
		for {
			select {
			case <-ctx.Done():
				statRepo.Add(stat)
				ctx, _ = context.WithTimeout(context.Background(), 60*time.Second)
			case <-ctxFinish.Done():
				return
			case stat = <-ch:
				// secondVolume += stat.Volume
				// if lastSec == -1 {
				// 	lastSec = stat.Time
				// 	tmpOhLCV.Open = stat.Price
				// 	tmpOhLCV.Ticker = text[0]
				// 	tmpOhLCV.Interval = 1
				// 	tmpOhLCV.Time = int32(sec)
				// 	secondMax = curPrice
				// 	secondMin = curPrice
				// 	secondVolume = volume
				// }
				// if sec != lastSec && lastSec != -1 {
				// 	tmpOhLCV.Close = lastPrice
				// 	tmpOhLCV.High = secondMax
				// 	tmpOhLCV.Low = secondMin
				// 	tmpOhLCV.Volume = uint32(secondVolume - volume)
				// 	answerList = append(answerList, tmpOhLCV)
				// 	secondMax = curPrice
				// 	secondMin = curPrice
				// 	tmpOhLCV.Time = int32(sec)
				// 	tmpOhLCV.Open = float32(val64)
				// 	secondVolume = volume
				// 	lastSec = sec
				// }
				// if curPrice > secondMax {
				// 	secondMax = curPrice
				// }
				// if curPrice < secondMin {
				// 	secondMin = curPrice
				// }
				// lastPrice = curPrice
			}
		}
	}()
	clentRepo := clients.NewRepository(db)
	orderHistoryRepo := order_history.NewRepository(db)
	requestRepo := requests.NewRepository(db)
	statiRepo := stats.NewRepository(db)

	sessions := session.NewSessionsManager(clentRepo)

	clintHandler := handlers.ClientHandler{
		Tmpl:             templates,
		ClientRepo:       clentRepo,
		OrderHistoryRepo: orderHistoryRepo,
		StatsRepo:        statiRepo,
		ReqestsRepo:      requestRepo,
		Sessions:         sessions,
	}

	r := mux.NewRouter()

	mux := middleware.Auth(sessions, r)
	mux = middleware.AccessLog(logger, mux)
	mux = middleware.Panic(mux)

	r.HandleFunc("/", clintHandler.Template).Methods("GET")
	r.HandleFunc("/api/v1/login", clintHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/register", clintHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/cancel", clintHandler.Cancel).Methods("POST")
	r.HandleFunc("/api/v1/deal", clintHandler.Deal).Methods("POST")
	r.HandleFunc("/api/v1/history", clintHandler.History).Methods("GET")
	r.HandleFunc("/api/v1/status", clintHandler.Stats).Methods("GET")

	http.ListenAndServe(":8080", r)
}
