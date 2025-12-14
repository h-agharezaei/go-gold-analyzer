# 🤝 راهنمای مشارکت در Gold Analyzer

خوشبخت‌ام که می‌خواهید در Gold Analyzer مشارکت کنید! این فایل نحوهٔ مشارکت در پروژه را توضیح می‌دهد.

## 📋 جدول محتویات

- [نحوهٔ شروع](#نحوهٔ-شروع)
- [فرآیند کار](#فرآیند-کار)
- [استانداردهای کد](#استانداردهای-کد)
- [Commit Messages](#commit-messages)
- [Pull Requests](#pull-requests)
- [Testing](#testing)
- [Documentation](#documentation)
- [رفتار و اخلاقیات](#رفتار-و-اخلاقیات)

---

## 🚀 نحوهٔ شروع

### 1. Fork کردن

```bash
# روی GitHub، روی دکمهٔ Fork کلیک کنید
# سپس:
git clone https://github.com/YOUR_USERNAME/gold-analyzer.git
cd gold-analyzer
git remote add upstream https://github.com/ORIGINAL_OWNER/gold-analyzer.git
```

### 2. ایجاد Branch

```bash
# بروزرسانی branch اصلی
git fetch upstream
git checkout main
git merge upstream/main

# ایجاد branch جدید برای feature/fix
git checkout -b feature/your-feature-name
# یا برای bug fix:
git checkout -b fix/bug-description
```

### 3. نصب وابستگی‌ها

```bash
go mod download
make install-tools
```

---

## 🔄 فرآیند کار

### مراحل توسعهٔ یک ویژگی

```
1. نوشتن کد
   ↓
2. اجرای تست‌ها (make test)
   ↓
3. فرمت کردن کد (make fmt)
   ↓
4. بررسی کیفیت (make lint)
   ↓
5. Commit و Push
   ↓
6. Pull Request
   ↓
7. Code Review
   ↓
8. Merge
```

### دستورات مفید

```bash
# تمام بررسی‌ها
make check

# فقط تست
make test

# فقط lint
make lint

# بنچمارک
make bench
```

---

## 📝 استانداردهای کد

### نام‌گذاری

```go
// ✅ درست
func CalculateRSI(closes []float64, period int) []float64 {}
type SignalType string
var maxRetries = 5

// ❌ اشتباه
func calculate_RSI(closes []float64, period int) {}
type signal_type string
var MAX_RETRIES = 5
```

### کامنت‌ها

```go
// ✅ درست - کامنت واضح
// CalculateRSI computes the Relative Strength Index
// for the given closing prices with the specified period.
func CalculateRSI(closes []float64, period int) []float64 {

// ✅ توضیح علت
if price < 0 {
    // Invalid price, skip this candle
    continue
}

// ❌ کامنت غیرضروری
// increment i
i++

// ❌ کامنت غلط
// this is a function
func Foo() {}
```

### Error Handling

```go
// ✅ درست
if err != nil {
    return nil, fmt.Errorf("failed to fetch data: %w", err)
}

// ❌ اشتباه
if err != nil {
    panic(err)
}
if err != nil {
    return nil, err  // بدون context
}
```

### طول خط

```
حداکثر 120 کاراکتر برای هر خط
```

---

## 💬 Commit Messages

### فرمت

```
[type]: brief description

optional body explaining why and what
```

### Types

```
feat:     new feature
fix:      bug fix
docs:     documentation
style:    formatting, missing semicolons, etc
refactor: code change that neither fixes a bug nor adds a feature
perf:     code change that improves performance
test:     adding or updating tests
chore:    changes to build process, dependencies, etc
ci:       changes to CI configuration
```

### مثال‌ها

```
✅ درست:
feat: add RSI calculation with customizable period
fix: handle nil values in MACD calculation
docs: update README with trading signals explanation
perf: optimize ATR calculation for large datasets

❌ اشتباه:
fixed stuff
updates
asdf
WIP
```

---

## 📤 Pull Requests

### قبل از Submit

- [ ] تمام تست‌ها pass می‌کنند: `make test`
- [ ] کد lint pass می‌کند: `make lint`
- [ ] کد فرمت شده است: `make fmt`
- [ ] CHANGELOG.md به‌روز شده است
- [ ] مستندات به‌روز شدهٔ است

### PR Title Format

```
[type]: brief description

Examples:
[feat]: add Bollinger Bands indicator
[fix]: resolve rate limiting issue
[docs]: improve QUICK_START guide
```

### PR Description Template

```markdown
## 📝 توضیح
توضیح کوتاه از تغییرات.

## 🎯 مرتبط با
- Closes #123
- Related to #456

## ✅ Checklist
- [ ] تست‌ها pass می‌کنند
- [ ] کد فرمت شده است
- [ ] مستندات به‌روز شده است
- [ ] CHANGELOG به‌روز شده است
- [ ] Breaking changes نیست

## 📸 Screenshot (اگر UI باشد)
```

---

## 🧪 Testing

### نوشتن تست

```go
func TestRSI(t *testing.T) {
    // Arrange: آماده‌سازی داده‌ها
    closes := []float64{44, 44.34, 44.09, ...}
    
    // Act: اجرای تابع
    result := indicators.RSI(closes, 14)
    
    // Assert: بررسی نتایج
    if len(result) != len(closes) {
        t.Errorf("Expected %d, got %d", len(closes), len(result))
    }
}
```

### Benchmark

```go
func BenchmarkRSI(b *testing.B) {
    closes := make([]float64, 1000)
    for i := 0; i < 1000; i++ {
        closes[i] = 44 + float64(i%10)
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        indicators.RSI(closes, 14)
    }
}
```

### اجرا

```bash
# تمام تست‌ها
go test ./...

# فقط یک تست
go test -run TestRSI ./...

# بنچمارک
go test -bench=. ./test

# Coverage
go test -cover ./...
```

---

## 📚 Documentation

### Docstrings

```go
// Package indicators provides technical analysis calculations
// for financial market data.
package indicators

// RSI calculates the Relative Strength Index for given closing prices.
// The RSI oscillates between 0 and 100, typically used to identify
// overbought (>70) and oversold (<30) conditions.
//
// Parameters:
//   - closes: slice of closing prices
//   - period: the lookback period (typically 14)
//
// Returns:
//   - slice of RSI values with the same length as closes
func RSI(closes []float64, period int) []float64 {
```

### README.md

```markdown
- نمایش صاف و واضح
- مثال‌های استفاده
- بخش Troubleshooting
- لینک‌های مفید
```

### CHANGELOG.md

```markdown
- هر release را بر روز کنید
- Semantic Versioning استفاده کنید
- مثال‌ها اضافه کنید
```

---

## 🐛 Bug Reports

### معلومات ضروری

```
## توصیف Bug
یک توضیح مختصر و واضح.

## نحوهٔ تکرار
مراحل دقیق برای تکرار مشکل.

## رفتار مورد انتظار
چه باید اتفاق بیفتد.

## رفتار واقعی
چه اتفاق افتاده است.

## Environment
- Go version: 1.25.2
- OS: macOS/Linux/Windows
- Branch: main/develop
```

---

## 💡 Feature Requests

### اطلاعات ضروری

```
## توضیح
توضیح ویژگی مورد نظر.

## مورد استفاده
چرا این ویژگی مهم است.

## راه‌حل پیشنهادی
اگر ایده‌ای دارید.

## جایگزین‌های ممکن
راه‌حل‌های دیگر.
```

---

## 🏆 بهترین روش‌ها

### 1. Commit کوچک و منطقی
```bash
# ✅ درست
git commit -m "feat: add RSI indicator"
git commit -m "test: add unit tests for RSI"

# ❌ اشتباه
git commit -m "fixed stuff and added features"
```

### 2. آپ‌دیت کردن از upstream
```bash
git fetch upstream
git rebase upstream/main
```

### 3. squash کردن commits قبل از PR
```bash
git rebase -i upstream/main
```

### 4. کامنت‌های واضح
```go
// ✅ خوب: علت را توضیح می‌دهد
// Skip zero values as they may cause division by zero
if value == 0 {
    continue
}

// ❌ بد: بیهوده
// Loop through values
for _, value := range values {
```

### 5. مستندات به‌روز
```markdown
- README.md رو اپدیت کنید
- CHANGELOG.md رو اپدیت کنید
- کامنت‌ها رو اپدیت کنید
```

---

## 📋 Checklist برای PR

```markdown
- [ ] کد با استانداردهای پروژه مطابقت دارد
- [ ] تست‌های جدید نوشته شدند
- [ ] تست‌های موجود pass می‌کنند
- [ ] کد فرمت شده است (make fmt)
- [ ] lint check pass می‌کند (make lint)
- [ ] مستندات به‌روز شدهٔ است
- [ ] CHANGELOG به‌روز شده است
- [ ] Commit messages واضح هستند
- [ ] هیچ breaking change نیست
- [ ] Performance test انجام شده است
```

---

## ❓ سؤالات متداول

### Q: چقدر طول می‌کشد code review؟
A: عادتاً 24-48 ساعت، اما می‌تواند بیشتر باشد.

### Q: اگر conflict داشته باشم؟
A: از git rebase استفاده کنید:
```bash
git fetch upstream
git rebase upstream/main
# resolve conflicts
git add .
git rebase --continue
git push --force-with-lease
```

### Q: می‌توانم چند PR باز کنم؟
A: بله! اما هر کدام برای یک ویژگی/bug باشد.

### Q: اگر PR رد شود؟
A: فکر نکنید بدی کردید! feedback سازنده است. سؤال کنید و اصلاح کنید.

---

## 🎓 منابع مفید

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://golang.org/doc/effective_go)
- [Keep a Changelog](https://keepachangelog.com/)
- [Semantic Versioning](https://semver.org/)

---

## 🙏 تشکر

تشکر بسیار از مشارکت شما! پروژه بدون کمک شما ممکن نیست.

---

## ✨ آخرین نکات

- بیش از همه، **مودب و احترام‌آمیز** باشید
- **صبور** باشید - code review زمان می‌برد
- **سؤال کنید** - اگر مطمئن نیستید
- **یاد بگیرید** - از feedback استفاده کنید
- **لذت ببرید** - توسعهٔ open source جالب است!

---

**Happy Contributing! 🎉**

**نسخهٔ آخری**: دسامبر 2025