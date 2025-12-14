package strategy

type Signal string

const (
	BUY  Signal = "BUY"
	SELL Signal = "SELL"
	HOLD Signal = "HOLD"
)

func GoldStrategy(
	rsi, macdHist, atr []float64,
	price float64,
) Signal {

	last := len(rsi) - 1

	// Buy Conditions
	if rsi[last] > 40 && rsi[last] < 55 &&
		macdHist[last] > 0 &&
		atr[last] > atr[last-1] {
		return BUY
	}

	// Sell Conditions
	if rsi[last] > 65 &&
		macdHist[last] < 0 {
		return SELL
	}

	return HOLD
}