package indicators

func ATR(high, low, close []float64, period int) []float64 {
	atr := make([]float64, len(close))
	tr := make([]float64, len(close))

	for i := 1; i < len(close); i++ {
		hL := high[i] - low[i]
		hC := abs(high[i] - close[i-1])
		lC := abs(low[i] - close[i-1])
		tr[i] = max(hL, max(hC, lC))
	}

	var sum float64
	for i := 1; i <= period; i++ {
		sum += tr[i]
	}
	atr[period] = sum / float64(period)

	for i := period + 1; i < len(tr); i++ {
		atr[i] = (atr[i-1]*float64(period-1) + tr[i]) / float64(period)
	}
	return atr
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
