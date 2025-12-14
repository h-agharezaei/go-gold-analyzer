package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gold-analyzer/config"
	"gold-analyzer/indicators"
	"gold-analyzer/shutdown"
	"gold-analyzer/strategy"
	"gold-analyzer/yahoo"
)

var lastSignal strategy.Signal

func main() {
	cfg := config.DefaultConfig()

	fmt.Println("🚀 Gold Analyzer - شروع نظارت خودکار...")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("⚙️  تنظیمات:\n")
	fmt.Printf("   • نماد: %s\n", cfg.Symbol)
	fmt.Printf("   • بازه زمانی: %s\n", cfg.Interval)
	fmt.Printf("   • محدوده: %s\n", cfg.Range)
	fmt.Printf("   • فاصله بررسی: %v\n", cfg.CheckInterval)
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("💡 برای متوقف کردن، Ctrl+C را فشار دهید...")

	// Create shutdown manager
	shutdownMgr := shutdown.NewManager()

	// Register shutdown hooks
	shutdownMgr.RegisterHook(func() error {
		return saveShutdownStats(cfg)
	})

	shutdownMgr.RegisterHook(func() error {
		return closeResources(cfg)
	})

	// Start signal handling
	shutdownMgr.Start()

	// Handle shutdown signal in a separate goroutine
	go func() {
		sig := shutdownMgr.WaitForShutdown()
		fmt.Printf("\n\n🛑 سیگنال دریافت شد: %v\n", sig)
		fmt.Println("⏳ درحال متوقف کردن برنامه...")
		shutdownMgr.Stop()
	}()

	// Create ticker for periodic checks
	ticker := time.NewTicker(cfg.CheckInterval)
	defer ticker.Stop()

	// اجرای اولی بدون تاخیر
	analyzeGold(cfg)

	// حلقه نظارت
	for {
		if !shutdownMgr.IsRunning() {
			// Perform graceful shutdown
			if err := shutdownMgr.Shutdown(cfg.ShutdownTimeout); err != nil {
				fmt.Printf("❌ خطا در طول Shutdown: %v\n", err)
			}
			shutdownMgr.SignalShutdownComplete()
			return
		}

		select {
		case <-ticker.C:
			if shutdownMgr.IsRunning() {
				analyzeGold(cfg)
			}

		case <-shutdownMgr.GetShutdownChan():
			return
		}
	}
}

// analyzeGold performs the gold analysis
func analyzeGold(cfg *config.Config) {
	now := time.Now()
	fmt.Printf("\n📊 بررسی در: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Println(strings.Repeat("-", 70))

	// دریافت داده‌ها
	candles, err := yahoo.FetchCandles(cfg.Symbol, cfg.Interval, cfg.Range)
	if err != nil {
		fmt.Printf("❌ خطا در دریافت داده: %v\n", err)
		logError(cfg, err.Error())
		return
	}

	if len(candles) == 0 {
		fmt.Println("⚠️  داده‌ای برای تجزیه و تحلیل وجود ندارد")
		return
	}

	// استخراج داده‌ها
	var close, high, low []float64
	for _, c := range candles {
		close = append(close, c.Close)
		high = append(high, c.High)
		low = append(low, c.Low)
	}

	// محاسبه اندیکاتورها
	rsi := indicators.RSI(close, cfg.RSIPeriod)
	macd, signal, hist := indicators.MACD(close)
	atr := indicators.ATR(high, low, close, cfg.ATRPeriod)

	lastIdx := len(close) - 1
	currentPrice := close[lastIdx]
	lastRSI := rsi[lastIdx]
	lastMACD := macd[lastIdx]
	lastSignalValue := signal[lastIdx]
	lastHist := hist[lastIdx]
	lastATR := atr[lastIdx]

	// نمایش قیمت فعلی
	fmt.Printf("\n💰 قیمت فعلی طلا: %.2f USD\n", currentPrice)

	// نمایش تغییر قیمت (اگر داده کافی باشد)
	if lastIdx > 0 {
		prevPrice := close[lastIdx-1]
		change := currentPrice - prevPrice
		changePercent := (change / prevPrice) * 100
		arrow := "↑"
		if change < 0 {
			arrow = "↓"
		}
		fmt.Printf("   %s تغییر: %.2f USD (%.2f%%)\n", arrow, change, changePercent)
	}

	// نمایش اندیکاتورها
	fmt.Println("\n📈 اندیکاتورهای تکنیکال:")
	fmt.Printf("   • RSI (%d):        %.2f", cfg.RSIPeriod, lastRSI)
	if lastRSI < cfg.RSIBuyLower {
		fmt.Print(" 🟦 فروش زیادی")
	} else if lastRSI > cfg.RSISellThreshold {
		fmt.Print(" 🟥 خرید زیادی")
	} else if lastRSI > cfg.RSIBuyLower && lastRSI < cfg.RSIBuyUpper {
		fmt.Print(" 🟩 محدوده مناسب")
	}
	fmt.Println()

	fmt.Printf("   • MACD:            %.6f\n", lastMACD)
	fmt.Printf("   • MACD Signal:     %.6f\n", lastSignalValue)
	fmt.Printf("   • MACD Histogram:  %.6f", lastHist)
	if lastHist > 0 {
		fmt.Print(" 📈 مثبت")
	} else {
		fmt.Print(" 📉 منفی")
	}
	fmt.Println()

	fmt.Printf("   • ATR (%d):        %.2f\n", cfg.ATRPeriod, lastATR)

	// محاسبه سیگنال
	strategySignal := strategy.GoldStrategy(rsi, hist, atr, currentPrice)

	// Store last signal
	lastSignal = strategySignal

	// نمایش سیگنال و توصیه
	fmt.Println("\n🎯 سیگنال معاملاتی:")
	switch strategySignal {
	case strategy.BUY:
		fmt.Println("   ✅ سیگنال: خریدش کن (BUY)")
		printBuyReason(cfg, lastRSI, lastHist, lastATR)
	case strategy.SELL:
		fmt.Println("   ❌ سیگنال: بفروش (SELL)")
		printSellReason(cfg, lastRSI, lastHist)
	case strategy.HOLD:
		fmt.Println("   ⏸️  سیگنال: نگاه کن (HOLD)")
		printHoldReason(lastRSI, lastHist, lastATR)
	}

	logSignal(cfg, strategySignal, currentPrice, lastRSI, lastHist, lastATR)
	fmt.Println(strings.Repeat("=", 70))
}

func printBuyReason(cfg *config.Config, rsi, hist, atr float64) {
	fmt.Println("   دلایل سیگنال خرید:")
	if rsi > cfg.RSIBuyLower && rsi < cfg.RSIBuyUpper {
		fmt.Printf("      • RSI در محدوده مناسب (%.2f - %.2f)\n", cfg.RSIBuyLower, cfg.RSIBuyUpper)
	}
	if hist > 0 {
		fmt.Println("      • MACD Histogram مثبت است")
	}
	fmt.Printf("      • ATR = %.2f (نوسان خوب)\n", atr)
}

func printSellReason(cfg *config.Config, rsi, hist float64) {
	fmt.Println("   دلایل سیگنال فروش:")
	if rsi > cfg.RSISellThreshold {
		fmt.Printf("      • RSI بالا است (%.2f > %.2f) - اشباع خریدار\n", rsi, cfg.RSISellThreshold)
	}
	if hist < 0 {
		fmt.Println("      • MACD Histogram منفی است - تضعیف روند")
	}
}

func printHoldReason(rsi, hist, atr float64) {
	fmt.Println("   دلایل سیگنال انتظار:")
	fmt.Printf("      • شرایط بازار مناسب برای ورود نیست\n")
	fmt.Printf("      • RSI = %.2f, MACD Hist = %.6f, ATR = %.2f\n", rsi, hist, atr)
}

func logSignal(cfg *config.Config, sig strategy.Signal, price, rsi, hist, atr float64) {
	if cfg.LogFile == "" {
		return
	}

	logEntry := fmt.Sprintf("[%s] Signal: %s | Price: %.2f | RSI: %.2f | MACD: %.6f | ATR: %.2f\n",
		time.Now().Format("2006-01-02 15:04:05"), sig, price, rsi, hist, atr)

	f, err := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(logEntry)
}

func logError(cfg *config.Config, errMsg string) {
	if cfg.LogFile == "" {
		return
	}

	logEntry := fmt.Sprintf("[%s] ERROR: %s\n", time.Now().Format("2006-01-02 15:04:05"), errMsg)

	f, err := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(logEntry)
}

// saveShutdownStats saves statistics before shutdown
func saveShutdownStats(cfg *config.Config) error {
	if cfg.LogFile == "" {
		return nil
	}

	logEntry := fmt.Sprintf("[%s] SHUTDOWN: برنامه با موفقیت متوقف شد - آخرین سیگنال: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), lastSignal)

	f, err := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(logEntry)
	if err == nil {
		fmt.Printf("   ✓ لاگ‌های نهایی ذخیره شدند (%s)\n", cfg.LogFile)
	}
	return err
}

// closeResources closes any open resources
func closeResources(cfg *config.Config) error {
	fmt.Println("\n🔐 بستن منابع...")
	fmt.Println("   ✓ تمام منابع بسته شدند")
	return nil
}
