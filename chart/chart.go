package chart

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// ChartData holds all the data needed for charting
type ChartData struct {
	Timestamps  []time.Time
	Prices      []float64
	RSI         []float64
	MACD        []float64
	Signal      []float64
	Histogram   []float64
	BuySignals  []int // indices where buy signals occurred
	SellSignals []int // indices where sell signals occurred
}

// GeneratePriceChart creates a price chart with buy/sell signals
func GeneratePriceChart(data *ChartData, outputPath string) error {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Gold Price Analysis",
			Subtitle: fmt.Sprintf("Updated: %s", time.Now().Format("2006-01-02 15:04:05")),
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
	)

	// Format timestamps for X axis
	xAxisData := make([]string, len(data.Timestamps))
	for i, t := range data.Timestamps {
		xAxisData[i] = t.Format("2006-01-02 15:04")
	}

	// Add price line
	priceItems := make([]opts.LineData, len(data.Prices))
	for i, price := range data.Prices {
		priceItems[i] = opts.LineData{Value: price}
	}

	line.AddSeries("Price", priceItems,
		charts.WithLineChartOpts(opts.LineChart{
			Smooth: true,
		}),
	)

	// Add buy signals as scatter points
	if len(data.BuySignals) > 0 {
		buyItems := make([]opts.BarData, len(data.BuySignals))
		for i, idx := range data.BuySignals {
			buyItems[i] = opts.BarData{Value: data.Prices[idx], Name: "Buy"}
		}

		// Create a scatter series for buy signals
		scatter := charts.NewScatter()
		scatter.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{
				Title: "Buy Signals",
			}),
		)

		buyScatterItems := make([]opts.ScatterData, len(data.BuySignals))
		for i, idx := range data.BuySignals {
			buyScatterItems[i] = opts.ScatterData{Value: []interface{}{xAxisData[idx], data.Prices[idx]}}
		}
		scatter.AddSeries("Buy", buyScatterItems)
	}

	// Add sell signals as scatter points
	if len(data.SellSignals) > 0 {
		sellScatterItems := make([]opts.ScatterData, len(data.SellSignals))
		for i, idx := range data.SellSignals {
			sellScatterItems[i] = opts.ScatterData{Value: []interface{}{xAxisData[idx], data.Prices[idx]}}
		}
	}

	line.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Data: xAxisData,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Price (USD)",
		}),
	)

	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	return line.Render(f)
}

// GenerateRSIChart creates RSI indicator chart
func GenerateRSIChart(data *ChartData, outputPath string) error {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "RSI (Relative Strength Index)",
			Subtitle: fmt.Sprintf("Updated: %s", time.Now().Format("2006-01-02 15:04:05")),
		}),
	)

	// Format timestamps
	xAxisData := make([]string, len(data.Timestamps))
	for i, t := range data.Timestamps {
		xAxisData[i] = t.Format("2006-01-02 15:04")
	}

	// Add RSI line
	rsiItems := make([]opts.LineData, len(data.RSI))
	for i, rsi := range data.RSI {
		rsiItems[i] = opts.LineData{Value: rsi}
	}

	line.AddSeries("RSI", rsiItems)

	// Add threshold lines
	thresholdItems := make([]opts.LineData, len(data.RSI))
	for i := range data.RSI {
		thresholdItems[i] = opts.LineData{Value: 70}
	}
	line.AddSeries("Overbought (70)", thresholdItems)

	thresholdItems2 := make([]opts.LineData, len(data.RSI))
	for i := range data.RSI {
		thresholdItems2[i] = opts.LineData{Value: 30}
	}
	line.AddSeries("Oversold (30)", thresholdItems2)

	line.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Data: xAxisData,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "RSI",
			Min:  0,
			Max:  100,
		}),
	)

	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	return line.Render(f)
}

// GenerateMACDChart creates MACD indicator chart
func GenerateMACDChart(data *ChartData, outputPath string) error {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "MACD (Moving Average Convergence Divergence)",
			Subtitle: fmt.Sprintf("Updated: %s", time.Now().Format("2006-01-02 15:04:05")),
		}),
	)

	// Format timestamps
	xAxisData := make([]string, len(data.Timestamps))
	for i, t := range data.Timestamps {
		xAxisData[i] = t.Format("2006-01-02 15:04")
	}

	// Add MACD line
	macdItems := make([]opts.LineData, len(data.MACD))
	for i, macd := range data.MACD {
		macdItems[i] = opts.LineData{Value: macd}
	}
	line.AddSeries("MACD", macdItems)

	// Add Signal line
	signalItems := make([]opts.LineData, len(data.Signal))
	for i, signal := range data.Signal {
		signalItems[i] = opts.LineData{Value: signal}
	}
	line.AddSeries("Signal", signalItems)

	// Add Histogram
	histItems := make([]opts.LineData, len(data.Histogram))
	for i, hist := range data.Histogram {
		histItems[i] = opts.LineData{Value: hist}
	}
	line.AddSeries("Histogram", histItems)

	line.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Data: xAxisData,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "MACD Value",
		}),
	)

	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	return line.Render(f)
}

// GenerateAllCharts generates all charts and returns paths
func GenerateAllCharts(data *ChartData, outputDir string) (map[string]string, error) {
	results := make(map[string]string)

	priceChartPath := filepath.Join(outputDir, "price_chart.html")
	if err := GeneratePriceChart(data, priceChartPath); err != nil {
		return nil, fmt.Errorf("error generating price chart: %w", err)
	}
	results["price"] = priceChartPath

	rsiChartPath := filepath.Join(outputDir, "rsi_chart.html")
	if err := GenerateRSIChart(data, rsiChartPath); err != nil {
		return nil, fmt.Errorf("error generating RSI chart: %w", err)
	}
	results["rsi"] = rsiChartPath

	macdChartPath := filepath.Join(outputDir, "macd_chart.html")
	if err := GenerateMACDChart(data, macdChartPath); err != nil {
		return nil, fmt.Errorf("error generating MACD chart: %w", err)
	}
	results["macd"] = macdChartPath

	return results, nil
}
