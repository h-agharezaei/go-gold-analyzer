package indicators

func EMA(data []float64, period int) []float64 {
	ema := make([]float64, len(data))
	k := 2.0 / float64(period+1)
	ema[0] = data[0]

	for i := 1; i < len(data); i++ {
		ema[i] = data[i]*k + ema[i-1]*(1-k)
	}
	return ema
}

func MACD(closes []float64) (macd, signal, hist []float64) {
	fast := EMA(closes, 8)
	slow := EMA(closes, 21)

	macd = make([]float64, len(closes))
	for i := range closes {
		macd[i] = fast[i] - slow[i]
	}

	signal = EMA(macd, 5)
	hist = make([]float64, len(macd))
	for i := range macd {
		hist[i] = macd[i] - signal[i]
	}
	return
}