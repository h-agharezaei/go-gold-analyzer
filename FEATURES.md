# 🎯 Gold Analyzer - Features Documentation

## 📊 Overview

Gold Analyzer is an automated technical analysis tool for gold (GC=F) that provides real-time signals and interactive visualizations.

---

## ✨ Core Features

### 1. 🔄 Automatic Monitoring
- **Auto-check every minute** (configurable)
- **Real-time data** from Yahoo Finance API
- **Non-blocking operation** - can run 24/7
- **Graceful shutdown** with proper resource cleanup

### 2. 📈 Technical Indicators

#### RSI (Relative Strength Index)
- **Period**: 14 days (configurable)
- **Purpose**: Identify overbought/oversold conditions
- **Thresholds**:
  - Below 30: Oversold (potential buy)
  - 30-70: Neutral zone
  - Above 70: Overbought (potential sell)
- **Usage**: Momentum analysis and reversal detection

#### MACD (Moving Average Convergence Divergence)
- **Fast Period**: 12 days
- **Slow Period**: 26 days
- **Signal Period**: 9 days
- **Components**:
  - MACD Line: Difference between fast and slow EMAs
  - Signal Line: 9-day EMA of MACD
  - Histogram: Difference between MACD and Signal
- **Usage**: Trend confirmation and momentum changes

#### ATR (Average True Range)
- **Period**: 14 days (configurable)
- **Purpose**: Measure market volatility
- **Usage**: Risk management and position sizing

### 3. 🎯 Trading Signals

The system generates three types of signals:

#### ✅ BUY Signal
**Conditions:**
- RSI between 40 and 55 (not overbought)
- MACD Histogram is positive
- ATR shows good volatility

**Meaning:** Good entry opportunity with manageable risk

#### ❌ SELL Signal
**Conditions:**
- RSI above 65 (overbought condition)
- MACD Histogram is negative

**Meaning:** Take profits or reduce exposure

#### ⏸️ HOLD Signal
**Conditions:**
- None of the above conditions are met

**Meaning:** Wait for better entry or exit opportunity

### 4. 📊 Interactive Charts & Dashboard

#### Generated Charts:
1. **Price Chart** (`price_chart.html`)
   - Historical gold prices
   - Buy signals marked with 🟢 triangles
   - Sell signals marked with 🔴 triangles
   - Interactive zoom and pan capabilities

2. **RSI Chart** (`rsi_chart.html`)
   - RSI indicator values
   - Overbought threshold (70) line
   - Oversold threshold (30) line
   - Color-coded zones

3. **MACD Chart** (`macd_chart.html`)
   - MACD line
   - Signal line
   - Histogram visualization
   - Trend change identification

#### Dashboard Features (`dashboard.html`):
- 🎨 **Responsive Design**: Works on desktop, tablet, and mobile
- 📱 **Modern UI**: Clean, professional interface with gradient backgrounds
- 🔗 **Navigation**: Quick links to specific charts
- 📚 **Educational Content**: Built-in explanations for each indicator
- ⚡ **Interactive**: Hover, zoom, pan, and download capabilities
- 🔄 **Real-time Updates**: Automatic timestamp updates
- ⚠️ **Disclaimers**: Clear risk warnings on dashboard

### 5. 📝 Signal Logging

**Key Feature:** Only BUY and SELL signals are logged (HOLD signals are filtered out)

#### Log Format:
```
[2025-12-16 22:41:35] Signal: BUY | Price: 4328.30 | RSI: 52.91 | MACD: 1.459146 | ATR: 29.50
[2025-12-16 23:41:35] Signal: SELL | Price: 4340.50 | RSI: 68.45 | MACD: -2.341000 | ATR: 31.20
```

#### Features:
- Timestamped entries
- Complete signal context (price, indicators)
- Searchable and analyzable
- Suitable for backtesting

### 6. ⚙️ Configuration Management

#### Default Settings:
```
Symbol:                 GC=F (Gold Futures)
Interval:              1h (Hourly data)
Range:                 7d (7 days of history)
Check Interval:        1 minute
RSI Period:            14
RSI Buy Lower:         40
RSI Buy Upper:         55
RSI Sell Threshold:    65
MACD Fast:             12
MACD Slow:             26
MACD Signal:           9
ATR Period:            14
```

#### Configuration Methods:

**Method 1: Environment Variables**
```bash
SYMBOL="GC=F" \
INTERVAL="1h" \
RANGE="7d" \
CHECK_INTERVAL_MINUTES="1" \
RSI_PERIOD="14" \
RSI_BUY_LOWER="40" \
RSI_BUY_UPPER="55" \
RSI_SELL_THRESHOLD="65" \
LOG_FILE="signals.log" \
CHART_OUTPUT_DIR="./charts" \
./analyzer
```

**Method 2: Command Line**
```bash
make run-with-charts  # Run with chart generation
make run-with-log     # Run with signal logging
make run              # Basic run
```

### 7. 🛑 Graceful Shutdown

#### Features:
- ✅ Catches SIGINT (Ctrl+C) and SIGTERM signals
- ✅ Completes ongoing analysis before shutdown
- ✅ Saves final signal state
- ✅ Closes file handles properly
- ✅ Releases all resources

#### Configuration:
```bash
SHUTDOWN_TIMEOUT_SECONDS=10 ./analyzer
```

---

## 📁 File Structure

```
gold-analyzer/
├── cmd/
│   └── main.go                 # Entry point and orchestration
├── chart/
│   ├── chart.go               # Chart generation (Price, RSI, MACD)
│   └── dashboard.go           # HTML dashboard creation
├── config/
│   └── config.go              # Configuration management
├── indicators/
│   ├── rsi.go                 # RSI calculation
│   ├── macd.go                # MACD calculation
│   ├── atr.go                 # ATR calculation
│   └── ema.go                 # EMA helper
├── model/
│   └── candle.go              # Candle data structure
├── strategy/
│   └── gold_strategy.go       # Signal generation logic
├── shutdown/
│   └── manager.go             # Graceful shutdown handling
├── yahoo/
│   └── client.go              # Yahoo Finance API client
├── test/
│   └── *.go                   # Unit tests
├── Makefile                   # Build and run commands
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
└── README.md                  # Quick start guide
```

---

## 🚀 Usage Examples

### Example 1: Basic Analysis
```bash
./analyzer
```
Output: Console display of current analysis only

### Example 2: With Logging
```bash
LOG_FILE="signals.log" ./analyzer
```
Output: Saves BUY/SELL signals to `signals.log`

### Example 3: With Charts
```bash
CHART_OUTPUT_DIR="./charts" ./analyzer
```
Output: Generates HTML charts in `./charts/` directory

### Example 4: Complete Setup
```bash
CHART_OUTPUT_DIR="./charts" \
LOG_FILE="signals.log" \
RSI_BUY_LOWER="35" \
RSI_SELL_THRESHOLD="70" \
./analyzer
```
Output: Custom thresholds + logging + charts

### Example 5: Using Makefile
```bash
# Build and run with all features
make dev-charts

# Just run with existing binary
make run-with-charts

# Run with logging only
make run-with-log
```

---

## 📊 Data Flow

```
┌─────────────────────┐
│  Yahoo Finance API  │
│   (GC=F futures)    │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│  Fetch Candles      │
│  (OHLCV data)       │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│  Calculate Indicators
│  • RSI              │
│  • MACD             │
│  • ATR              │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│  Generate Signals   │
│  • BUY              │
│  • SELL             │
│  • HOLD             │
└──────────┬──────────┘
           │
        ┌──┴──┐
        │     │
        ▼     ▼
    ┌─────┐ ┌──────────┐
    │ Log │ │  Charts  │
    └─────┘ └──────────┘
        │        │
        ▼        ▼
    [File]   [HTML Files]
```

---

## 🔧 Technical Details

### API Used
- **Yahoo Finance**: Free API, no authentication required
- **Retry Logic**: Up to 5 attempts with exponential backoff
- **Rate Limiting**: Handles 429 responses gracefully
- **Timeout**: 10 seconds per request

### Chart Library
- **go-echarts/v2**: Modern charting library
- **ECharts 5**: JavaScript rendering engine
- **Responsive**: Works on all screen sizes
- **Interactive**: Zoom, pan, save features built-in

### Performance
- **Memory Efficient**: Stores only necessary data
- **CPU Light**: Single-threaded analysis
- **Network Efficient**: Minimal API calls
- **Storage**: Charts are HTML files (~10-30 KB each)

---

## 📈 Indicator Formulas

### RSI (Relative Strength Index)
```
RS = Average Gain / Average Loss
RSI = 100 - (100 / (1 + RS))
```

### MACD
```
MACD = EMA(12) - EMA(26)
Signal = EMA(MACD, 9)
Histogram = MACD - Signal
```

### ATR (Average True Range)
```
TR = max(High-Low, abs(High-Close[prev]), abs(Low-Close[prev]))
ATR = SMA(TR, 14)
```

---

## ⚠️ Important Notes

### Accuracy
- Indicators are calculated on hourly data
- Results are advisory only
- Market conditions change rapidly
- Past performance ≠ Future results

### Risk Management
- Always use proper position sizing
- Set stop-loss orders
- Don't risk more than 2% per trade
- Use take-profit targets

### Disclaimer
Trading involves substantial risk. This tool is for educational and analytical purposes. Not financial advice. Always consult with qualified professionals before trading.

---

## 🔐 Security & Privacy

- ✅ No API keys required
- ✅ No personal data collected
- ✅ No database connections
- ✅ Local computation only
- ✅ Read-only API access
- ✅ HTTPS for data transfer

---

## 🐛 Known Limitations

1. **Historical Data Only**: Uses past 7 days of hourly data
2. **Single Symbol**: Configured for gold futures (GC=F)
3. **No Multi-timeframe**: Single 1h timeframe
4. **No Backtesting**: Real-time analysis only
5. **No Notifications**: Console and file output only
6. **Manual Trading**: No automated execution

---

## 🚀 Future Enhancements

Planned features for future versions:
- [ ] Email/SMS notifications
- [ ] Multiple symbol support
- [ ] Backtesting framework
- [ ] Web dashboard with real-time updates
- [ ] Telegram/Discord bot integration
- [ ] Database logging (SQLite, PostgreSQL)
- [ ] Advanced charting (Plotly, custom indicators)
- [ ] Position management
- [ ] Performance tracking
- [ ] Machine learning signals

---

**Version**: 1.0.0  
**Last Updated**: December 2025  
**License**: MIT