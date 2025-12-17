package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	// Symbol to analyze
	Symbol string
	// Interval for fetching data (1m, 5m, 1h, 1d, etc.)
	Interval string
	// Range for fetching data (1d, 5d, 1mo, 3mo, 6mo, 1y, 2y, 5y, 10y, ytd, max)
	Range string
	// Check interval in minutes
	CheckInterval time.Duration
	// RSI Period
	RSIPeriod int
	// MACD Fast Period
	MACDFastPeriod int
	// MACD Slow Period
	MACDSlowPeriod int
	// MACD Signal Period
	MACDSignalPeriod int
	// ATR Period
	ATRPeriod int
	// RSI Buy threshold (lower bound)
	RSIBuyLower float64
	// RSI Buy threshold (upper bound)
	RSIBuyUpper float64
	// RSI Sell threshold
	RSISellThreshold float64
	// Enable notifications
	EnableNotifications bool
	// Log file path (empty to disable)
	LogFile string
	// Chart output directory
	ChartOutputDir string
	// Shutdown timeout duration
	ShutdownTimeout time.Duration
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	cfg := &Config{
		Symbol:              "GC=F",
		Interval:            "1h",
		Range:               "7d",
		CheckInterval:       1 * time.Minute,
		RSIPeriod:           14,
		MACDFastPeriod:      12,
		MACDSlowPeriod:      26,
		MACDSignalPeriod:    9,
		ATRPeriod:           14,
		RSIBuyLower:         40,
		RSIBuyUpper:         55,
		RSISellThreshold:    65,
		EnableNotifications: false,
		LogFile:             "",
		ChartOutputDir:      "./charts",
		ShutdownTimeout:     5 * time.Second,
	}

	// Override with environment variables if present
	loadFromEnv(cfg)

	return cfg
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(cfg *Config) {
	if symbol := os.Getenv("SYMBOL"); symbol != "" {
		cfg.Symbol = symbol
	}
	if interval := os.Getenv("INTERVAL"); interval != "" {
		cfg.Interval = interval
	}
	if rangeVal := os.Getenv("RANGE"); rangeVal != "" {
		cfg.Range = rangeVal
	}
	if checkInterval := os.Getenv("CHECK_INTERVAL_MINUTES"); checkInterval != "" {
		if minutes, err := strconv.Atoi(checkInterval); err == nil {
			cfg.CheckInterval = time.Duration(minutes) * time.Minute
		}
	}
	if rsiPeriod := os.Getenv("RSI_PERIOD"); rsiPeriod != "" {
		if period, err := strconv.Atoi(rsiPeriod); err == nil {
			cfg.RSIPeriod = period
		}
	}
	if atrPeriod := os.Getenv("ATR_PERIOD"); atrPeriod != "" {
		if period, err := strconv.Atoi(atrPeriod); err == nil {
			cfg.ATRPeriod = period
		}
	}
	if rsiBuyLower := os.Getenv("RSI_BUY_LOWER"); rsiBuyLower != "" {
		if val, err := strconv.ParseFloat(rsiBuyLower, 64); err == nil {
			cfg.RSIBuyLower = val
		}
	}
	if rsiBuyUpper := os.Getenv("RSI_BUY_UPPER"); rsiBuyUpper != "" {
		if val, err := strconv.ParseFloat(rsiBuyUpper, 64); err == nil {
			cfg.RSIBuyUpper = val
		}
	}
	if rsiSellThreshold := os.Getenv("RSI_SELL_THRESHOLD"); rsiSellThreshold != "" {
		if val, err := strconv.ParseFloat(rsiSellThreshold, 64); err == nil {
			cfg.RSISellThreshold = val
		}
	}
	if logFile := os.Getenv("LOG_FILE"); logFile != "" {
		cfg.LogFile = logFile
	}
	if chartOutputDir := os.Getenv("CHART_OUTPUT_DIR"); chartOutputDir != "" {
		cfg.ChartOutputDir = chartOutputDir
	}
	if enableNotif := os.Getenv("ENABLE_NOTIFICATIONS"); enableNotif != "" {
		cfg.EnableNotifications = enableNotif == "1" || enableNotif == "true"
	}
	if shutdownTimeout := os.Getenv("SHUTDOWN_TIMEOUT_SECONDS"); shutdownTimeout != "" {
		if seconds, err := strconv.Atoi(shutdownTimeout); err == nil {
			cfg.ShutdownTimeout = time.Duration(seconds) * time.Second
		}
	}
}
