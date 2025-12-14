package main

import (
	"fmt"

	"gold-analyzer/indicators"
	"gold-analyzer/strategy"
	"gold-analyzer/yahoo"
)

func main() {
	candles, err := yahoo.FetchCandles("XAUUSD=X", "1h", "7d")
	if err != nil {
		panic(err)
	}

	var close, high, low []float64
	for _, c := range candles {
		close = append(close, c.Close)
		high = append(high, c.High)
		low = append(low, c.Low)
	}

	rsi := indicators.RSI(close, 14)
	_, _, hist := indicators.MACD(close)
	atr := indicators.ATR(high, low, close, 14)

	signal := strategy.GoldStrategy(
		rsi,
		hist,
		atr,
		close[len(close)-1],
	)

	fmt.Println("Gold Signal:", signal)
}