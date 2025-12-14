package yahoo

import (
	"encoding/json"
	"net/http"

	"gold-analyzer/model"
)

type response struct {
	Chart struct {
		Result []struct {
			Timestamp []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Open   []float64 `json:"open"`
					High   []float64 `json:"high"`
					Low    []float64 `json:"low"`
					Close  []float64 `json:"close"`
					Volume []int64   `json:"volume"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

func FetchCandles(symbol, interval, rangeVal string) ([]model.Candle, error) {
	url := "https://query1.finance.yahoo.com/v8/finance/chart/" +
		symbol + "?interval=" + interval + "&range=" + rangeVal

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	r := res.Chart.Result[0]
	q := r.Indicators.Quote[0]

	var candles []model.Candle
	for i := range r.Timestamp {
		candles = append(candles, model.Candle{
			Time:   r.Timestamp[i],
			Open:   q.Open[i],
			High:   q.High[i],
			Low:    q.Low[i],
			Close:  q.Close[i],
			Volume: q.Volume[i],
		})
	}
	return candles, nil
}