package utils

import (
	"bufio"
	"exchange/exchange/pkg/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filePath string) ([]service.OHLCV, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	answerList := make([]service.OHLCV, 0, 100)
	lastSec := -1
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	tmpOhLCV := service.OHLCV{}
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
