package yahoo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gold-analyzer/model"
)

type response struct {
	Chart struct {
		Result []struct {
			Timestamp  []int64 `json:"timestamp"`
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

	var res response
	var lastErr error

	// Retry logic with exponential backoff
	for attempt := 0; attempt < 5; attempt++ {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		// Add User-Agent to avoid 429 errors
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			waitTime := time.Duration(1<<uint(attempt)) * time.Second
			fmt.Printf("Attempt %d failed, waiting %v before retry...\n", attempt+1, waitTime)
			time.Sleep(waitTime)
			continue
		}
		defer resp.Body.Close()

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// Check for rate limiting
		if resp.StatusCode == 429 {
			waitTime := time.Duration(1<<uint(attempt)) * 2 * time.Second
			fmt.Printf("Rate limited (429), waiting %v before retry...\n", waitTime)
			time.Sleep(waitTime)
			continue
		}

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
		}

		// Try to parse JSON
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %w", err)
		}

		// If successful, break out of retry loop
		if len(res.Chart.Result) > 0 && len(res.Chart.Result[0].Indicators.Quote) > 0 {
			break
		}

		lastErr = fmt.Errorf("invalid response structure")
	}

	// Check if we got valid data
	if len(res.Chart.Result) == 0 {
		return nil, fmt.Errorf("no data in response after retries: %v", lastErr)
	}

	r := res.Chart.Result[0]
	if len(r.Indicators.Quote) == 0 {
		return nil, fmt.Errorf("no quote data in response")
	}

	q := r.Indicators.Quote[0]

	var candles []model.Candle
	for i := range r.Timestamp {
		// Skip entries with null values
		if i >= len(q.Open) || i >= len(q.High) || i >= len(q.Low) || i >= len(q.Close) || i >= len(q.Volume) {
			continue
		}
		if q.Open[i] == 0 && q.High[i] == 0 && q.Low[i] == 0 && q.Close[i] == 0 {
			continue
		}

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
