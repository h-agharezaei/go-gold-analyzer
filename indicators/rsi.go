package indicators

func RSI(closes []float64, period int) []float64 {
	rsi := make([]float64, len(closes))

	var gain, loss float64
	for i := 1; i <= period; i++ {
		diff := closes[i] - closes[i-1]
		if diff > 0 {
			gain += diff
		} else {
			loss -= diff
		}
	}

	avgGain := gain / float64(period)
	avgLoss := loss / float64(period)

	for i := period + 1; i < len(closes); i++ {
		diff := closes[i] - closes[i-1]
		var g, l float64
		if diff > 0 {
			g = diff
		} else {
			l = -diff
		}
		avgGain = (avgGain*float64(period-1) + g) / float64(period)
		avgLoss = (avgLoss*float64(period-1) + l) / float64(period)

		if avgLoss == 0 {
			rsi[i] = 100
		} else {
			rsi[i] = 100 - (100 / (1 + avgGain/avgLoss))
		}
	}
	return rsi
}
