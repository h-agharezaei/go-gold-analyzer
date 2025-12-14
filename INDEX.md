# 📑 فهرست Gold Analyzer

## 🎯 شروع سریع
- **[QUICK_START.md](./QUICK_START.md)** - راهنمای 5 دقیقه‌ای برای شروع فوری
- **[README.md](./README.md)** - مستندات کامل پروژه

## 📚 مستندات تفصیلی
- **[PROJECT_STATUS.md](./PROJECT_STATUS.md)** - وضعیت پروژه و آمار
- **[CHANGELOG.md](./CHANGELOG.md)** - تاریخچهٔ تغییرات و نقشهٔ راه
- **[CONTRIBUTING.md](./CONTRIBUTING.md)** - راهنمای مشارکت در پروژه

## 🗂️ ساختار کد

### نقطهٔ ورود
- **[cmd/main.go](./cmd/main.go)** - برنامهٔ اصلی
  - نظارت خودکار هر یک دقیقه
  - نمایش سیگنال‌ها و اندیکاتورها
  - Logging و error handling

### تنظیمات
- **[config/config.go](./config/config.go)** - سیستم تنظیمات
  - تنظیمات پیش‌فرض
  - Load from environment variables
  - تغییرات قابل سفارشی

### اندیکاتورهای تکنیکال
- **[indicators/rsi.go](./indicators/rsi.go)** - RSI Indicator
- **[indicators/macd.go](./indicators/macd.go)** - MACD Indicator
- **[indicators/atr.go](./indicators/atr.go)** - ATR Indicator

### منطق معاملاتی
- **[strategy/gold_strategy.go](./strategy/gold_strategy.go)** - سیگنال‌های معاملاتی
  - BUY، SELL، HOLD logic
  - شرایط سیگنال‌ها

### داده‌ها
- **[model/candle.go](./model/candle.go)** - ساختار داده Candle

### API Integration
- **[yahoo/client.go](./yahoo/client.go)** - Yahoo Finance Client
  - دریافت داده‌ها
  - Retry logic
  - Rate limiting handling

### تست‌ها
- **[test/indicators_test.go](./test/indicators_test.go)** - تست‌ها و Benchmarks
  - Unit tests
  - Benchmark tests

## 🛠️ ابزارهای DevOps
- **[Makefile](./Makefile)** - دستورات خودکار (20+ command)
- **[Dockerfile](./Dockerfile)** - Docker image
- **[docker-compose.yml](./docker-compose.yml)** - Docker Compose
- **[.env.example](./.env.example)** - نمونهٔ تنظیمات
- **[.gitignore](./.gitignore)** - Git exclusions

## 📋 فایل‌های پیکربندی
- **[go.mod](./go.mod)** - Go module definition

---

## 🚀 دستورات سریع

### اجرا
```bash
# اجرای ساده
go run ./cmd/main.go

# با logging
LOG_FILE="signals.log" go run ./cmd/main.go

# برای Docker
docker-compose up -d
```

### توسعهٔ
```bash
# تمام بررسی‌ها
make check

# تست‌ها
make test

# بنچمارک
make bench

# فرمت و lint
make fmt lint

# ساخت binary
make build
```

---

## 🎯 نمادهای محبوب
| نماد | توضیح |
|------|-------|
| `GC=F` | طلا (Futures) |
| `EURUSD=X` | یورو/دلار |
| `ES=F` | S&P 500 |
| `CL=F` | نفت خام |

---

## 📊 اندیکاتورها

### RSI (Relative Strength Index)
- **فایل**: `indicators/rsi.go`
- **دوره پیش‌فرض**: 14
- **محدودهٔ**: 0-100
- **استفاده**: شناسایی اشباع خریدار (>70) و فروش (>30)

### MACD (Moving Average Convergence Divergence)
- **فایل**: `indicators/macd.go`
- **دوره‌های**: 12، 26، 9
- **استفاده**: شناسایی روند و momentum

### ATR (Average True Range)
- **فایل**: `indicators/atr.go`
- **دوره پیش‌فرض**: 14
- **استفاده**: اندازهٔ نوسانات

---

## 🎯 سیگنال‌ها

### BUY ✅
شرایط:
- RSI بین 40-55
- MACD Histogram > 0
- ATR صعودی

### SELL ❌
شرایط:
- RSI > 65
- MACD Histogram < 0

### HOLD ⏸️
شرایط:
- سایر حالات

---

## 📈 Performance
```
RSI:    5,289 ns/op   (246K ops)
ATR:    8,892 ns/op   (120K ops)
MACD:  17,512 ns/op    (69K ops)
```

---

## 🔧 نحوهٔ تغییر تنظیمات

### از طریق Environment Variables
```bash
SYMBOL=EURUSD=X \
INTERVAL=1h \
RANGE=7d \
CHECK_INTERVAL_MINUTES=5 \
RSI_PERIOD=21 \
go run ./cmd/main.go
```

### از طریق .env
```bash
cp .env.example .env
# ویرایش .env
go run ./cmd/main.go
```

---

## 🆘 مشکلات متداول

### "Symbol may be delisted"
→ نماد را تصحیح کنید

### Rate Limiting (429)
→ برنامه خودکار retry می‌کند

### Logging نمی‌کند
→ اجازهٔ نوشتن در دایرکتوری را بررسی کنید

---

## 📞 منابع مفید
- [Go Documentation](https://golang.org/doc)
- [Yahoo Finance API](https://finance.yahoo.com)
- [Technical Analysis](https://en.wikipedia.org/wiki/Technical_analysis)

---

## ✨ نکات مهم
1. **محصول Production**: نسخهٔ 1.0.0 کامل است
2. **تست‌شده**: تمام اندیکاتورها تست‌شده‌اند
3. **قابل سفارشی**: تمام پارامترها قابل تغییر
4. **Docker Ready**: آماده برای containerization
5. **Well Documented**: مستندات کامل

---

## 🎓 راهنمای قدم‌به‌قدم

1. **بخش اول**: [QUICK_START.md](./QUICK_START.md) را بخوانید
2. **بخش دوم**: `go run ./cmd/main.go` اجرا کنید
3. **بخش سوم**: سیگنال‌ها را مراقب کنید
4. **بخش چهارم**: تنظیمات را برحسب نیاز تغییر دهید

---

**برای شروع، به [QUICK_START.md](./QUICK_START.md) بروید! 🚀**