package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OHLCV struct {
	ID       int64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Time     int32   `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	Interval int32   `protobuf:"varint,3,opt,name=Interval,proto3" json:"Interval,omitempty"`
	Open     float32 `protobuf:"fixed32,4,opt,name=Open,proto3" json:"Open,omitempty"`
	High     float32 `protobuf:"fixed32,5,opt,name=High,proto3" json:"High,omitempty"`
	Low      float32 `protobuf:"fixed32,6,opt,name=Low,proto3" json:"Low,omitempty"`
	Close    float32 `protobuf:"fixed32,7,opt,name=Close,proto3" json:"Close,omitempty"`
	Volume   uint32  `protobuf:"varint,8,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Ticker   string  `protobuf:"bytes,9,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
}

func parseFile(filePath string) ([]OHLCV, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	answerList := make([]OHLCV, 0, 100)
	lastSec := -1
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	tmpOhLCV := OHLCV{}
	var lastPrice float32
	var secondMax float32
	var secondMin float32
	var curPrice float32
	var secondVolume int
	var volume int
	var allVolume uint32
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		sec, _ := strconv.Atoi(text[3])
		val64, _ := strconv.ParseFloat(text[4], 1)
		curPrice = float32(val64)
		volume, _ = strconv.Atoi(text[5])
		allVolume += uint32(volume)
		secondVolume += volume
		if lastSec == -1 {
			lastSec = sec
			tmpOhLCV.Open = curPrice
			tmpOhLCV.Ticker = text[0]
			tmpOhLCV.Interval = 1
			tmpOhLCV.Time = int32(sec)
			secondMax = curPrice
			secondMin = curPrice
			secondVolume = volume
		}
		if sec != lastSec && lastSec != -1 {
			tmpOhLCV.Close = lastPrice
			tmpOhLCV.High = secondMax
			tmpOhLCV.Low = secondMin
			tmpOhLCV.Volume = uint32(secondVolume - volume)
			answerList = append(answerList, tmpOhLCV)
			secondMax = curPrice
			secondMin = curPrice
			tmpOhLCV.Time = int32(sec)
			tmpOhLCV.Open = float32(val64)
			secondVolume = volume
			lastSec = sec
		}
		if curPrice > secondMax {
			secondMax = curPrice
		}
		if curPrice < secondMin {
			secondMin = curPrice
		}
		lastPrice = curPrice
	}
	tmpOhLCV.Close = lastPrice
	tmpOhLCV.High = secondMax
	tmpOhLCV.Low = secondMin
	tmpOhLCV.Volume = uint32(secondVolume)
	answerList = append(answerList, tmpOhLCV)
	// var tmp uint32
	// for _, elem := range answerList {
	// 	tmp += elem.Volume
	// }
	// fmt.Println(tmp, allVolume)
	// fmt.Println(answerList[0].Volume)
	// fmt.Println(answerList[len(answerList)-2].Time)
	// for _, elem := range answerList {
	// 	fmt.Println(elem.Time)
	// }
	return answerList, nil
}

func main() {
	parseFile("SPFB.RTS_190517_190517.txt")
}
