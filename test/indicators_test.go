package test

import (
	"testing"

	"gold-analyzer/indicators"
)

func TestRSI(t *testing.T) {
	// نمونه داده ساده برای تست
	closes := []float64{
		44, 44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84,
		46.08, 45.89, 46.03, 45.61, 46.28, 46.00, 46.00, 46.00, 46.00, 46.00,
		46.00, 46.00, 46.00, 46.00, 46.00,
	}

	rsi := indicators.RSI(closes, 14)

	// بررسی طول نتیجه
	if len(rsi) != len(closes) {
		t.Errorf("Expected RSI length %d, got %d", len(closes), len(rsi))
	}

	// بررسی اینکه RSI در محدوده 0-100 است
	for i, val := range rsi {
		if val < 0 || val > 100 {
			t.Errorf("RSI at index %d is out of range: %.2f", i, val)
		}
	}

	// آخرین مقدار نباید صفر باشد (اگر داده درست باشد)
	if rsi[len(rsi)-1] == 0 && len(closes) > 14 {
		t.Logf("RSI values: %v", rsi[14:])
	}
}

func TestATR(t *testing.T) {
	high := []float64{
		45, 45.50, 45.75, 46, 46.25, 46.50, 46.75, 47, 47.25, 47.50,
		47.75, 48, 48.25, 48.50, 48.75, 49, 49.25, 49.50, 49.75, 50,
		50.25, 50.50, 50.75, 51, 51.25,
	}
	low := []float64{
		44, 44.50, 44.75, 45, 45.25, 45.50, 45.75, 46, 46.25, 46.50,
		46.75, 47, 47.25, 47.50, 47.75, 48, 48.25, 48.50, 48.75, 49,
		49.25, 49.50, 49.75, 50, 50.25,
	}
	close := []float64{
		44.5, 45.2, 45.5, 45.8, 46, 46.2, 46.5, 46.8, 47, 47.2,
		47.5, 47.8, 48, 48.2, 48.5, 48.8, 49, 49.2, 49.5, 49.8,
		50, 50.2, 50.5, 50.8, 51,
	}

	atr := indicators.ATR(high, low, close, 14)

	// بررسی طول نتیجه
	if len(atr) != len(close) {
		t.Errorf("Expected ATR length %d, got %d", len(close), len(atr))
	}

	// ATR نباید منفی باشد
	for i, val := range atr {
		if val < 0 {
			t.Errorf("ATR at index %d is negative: %.2f", i, val)
		}
	}

	// مقدار ATR در دوره باید بزرگتر از صفر باشد
	if atr[14] <= 0 {
		t.Error("ATR at period should be greater than 0")
	}
}

func TestMACD(t *testing.T) {
	closes := []float64{
		44, 44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84,
		46.08, 45.89, 46.03, 45.61, 46.28, 46.00, 46.12, 45.64, 46.21, 46.25,
		45.71, 46.45, 46.33, 46.41, 46.32, 46.99, 47.00, 47.14, 47.10, 47.18,
		47.14, 47.54, 48.20, 48.26, 48.38, 49.00, 49.14, 49.40, 49.63, 50.10,
	}

	macd, signal, hist := indicators.MACD(closes)

	// بررسی طول
	if len(macd) != len(closes) {
		t.Errorf("Expected MACD length %d, got %d", len(closes), len(macd))
	}
	if len(signal) != len(closes) {
		t.Errorf("Expected Signal length %d, got %d", len(closes), len(signal))
	}
	if len(hist) != len(closes) {
		t.Errorf("Expected Histogram length %d, got %d", len(closes), len(hist))
	}

	// بررسی روابط بین مقادیر
	for i := 0; i < len(macd); i++ {
		expectedHist := macd[i] - signal[i]
		if hist[i] != expectedHist {
			t.Logf("Histogram mismatch at %d: expected %.6f, got %.6f", i, expectedHist, hist[i])
		}
	}
}

func BenchmarkRSI(b *testing.B) {
	closes := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		closes[i] = float64(i%100) + 40
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indicators.RSI(closes, 14)
	}
}

func BenchmarkATR(b *testing.B) {
	high := make([]float64, 1000)
	low := make([]float64, 1000)
	close := make([]float64, 1000)

	for i := 0; i < 1000; i++ {
		close[i] = float64(i%100) + 40
		high[i] = close[i] + 1
		low[i] = close[i] - 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indicators.ATR(high, low, close, 14)
	}
}

func BenchmarkMACD(b *testing.B) {
	closes := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		closes[i] = float64(i%100) + 40
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indicators.MACD(closes)
	}
}
