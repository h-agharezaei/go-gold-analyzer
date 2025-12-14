# 🛑 Graceful Shutdown Guide

## مقدمه

Graceful shutdown یک روند نرم برای خاتمهٔ برنامه است که:
- ✅ عملیات جاری را تکمیل می‌کند
- ✅ منابع را به درستی بسته می‌کند
- ✅ داده‌ها را ذخیره می‌کند
- ✅ سیگنال‌های سیستمی را پردازش می‌کند

---

## 🎯 ویژگی‌های Graceful Shutdown

### سیگنال‌های پشتیبانی‌شده
- **SIGINT** (Ctrl+C) - درخواست توقف از کاربر
- **SIGTERM** - سیگنال خاتمهٔ از سیستم (Docker, systemd, etc)
- **SIGHUP** - سیگنال آپ‌دیت تنظیمات

### عملیات Shutdown
```
1. دریافت سیگنال
   ↓
2. متوقف کردن قبول کارهای جدید
   ↓
3. انتظار برای تکمیل عملیات جاری
   ↓
4. اجرای shutdown hooks
   ↓
5. بستن منابع
   ↓
6. ذخیرهٔ آمار نهایی
   ↓
7. خروج صحیح
```

---

## 📝 نحوهٔ استفاده

### اجرای برنامه
```bash
./analyzer
```

### متوقف کردن برنامه
```bash
# روش 1: Ctrl+C
Ctrl+C

# روش 2: سیگنال SIGTERM (از طریق kill)
kill -TERM <PID>

# روش 3: سیگنال SIGINT
kill -INT <PID>
```

### مثال پیاده‌سازی

```go
// Graceful shutdown فعال است در:
// cmd/main.go → استفاده از shutdown.Manager
// shutdown/manager.go → منطق مدیریت شدن

// شما می‌توانید shutdown hooks اضافه کنید:
shutdownMgr.RegisterHook(func() error {
    // کد تمیز‌کاری شما اینجا
    return nil
})
```

---

## 🔧 تنظیمات Shutdown

### Timeout
مدت زمانی که برنامه برای تکمیل عملیات صبر می‌کند:

```bash
# تنظیم از طریق environment variable
SHUTDOWN_TIMEOUT_SECONDS=10 ./analyzer

# یا از طریق .env
SHUTDOWN_TIMEOUT_SECONDS=10
```

### پیش‌فرض
- **5 ثانیه** - مدت زمان پیش‌فرض برای shutdown

---

## 📊 خروجی Shutdown

هنگام فشار دادن Ctrl+C یا ارسال سیگنال:

```
🛑 سیگنال دریافت شد: interrupt
⏳ درحال متوقف کردن برنامه...

======================================================================
🔄 شروع خاتمهٔ نرم برنامه...
======================================================================

🔐 بستن منابع...
   ✓ لاگ‌های نهایی ذخیره شدند
   ✓ تمام منابع بسته شدند

📊 آمار نهایی:
   ✓ تمام hooks تکمیل شدند
   ✓ مدت زمان Shutdown: 1.234s
   ✓ زمان خاتمه: 2025-12-14 22:53:14

======================================================================
✅ برنامه با موفقیت متوقف شد
======================================================================
```

---

## 🏗️ معماری Shutdown

### Shutdown Manager

```
Package: shutdown/manager.go

سازمان‌دهی:
- NewManager()              → ایجاد مدیریت
- Start()                   → شروع signal handling
- RegisterHook()            → ثبت تابع تمیز‌کاری
- Stop()                    → توقف برنامه
- Shutdown()                → اجرای تمیز‌کاری
- IsRunning()               → بررسی وضعیت
- WaitForShutdown()         → انتظار برای سیگنال
```

### Main Loop

```go
for {
    if !shutdownMgr.IsRunning() {
        shutdownMgr.Shutdown(cfg.ShutdownTimeout)
        return
    }
    
    select {
    case <-ticker.C:
        analyzeGold(cfg)
    }
}
```

---

## 🔐 Shutdown Hooks

تابع‌هایی که هنگام shutdown اجرا می‌شوند:

### Hook 1: ذخیرهٔ آمار
```go
shutdownMgr.RegisterHook(func() error {
    return saveShutdownStats(cfg)
})
```

### Hook 2: بستن منابع
```go
shutdownMgr.RegisterHook(func() error {
    return closeResources(cfg)
})
```

### Hook سفارشی
```go
shutdownMgr.RegisterHook(func() error {
    // کد خود را اینجا بنویسید
    fmt.Println("درحال بستن پایگاه داده...")
    // db.Close()
    return nil
})
```

---

## ⚡ حالات مختلف

### حالت عادی (Graceful)
```
اجرا → بررسی می‌کند → Ctrl+C → Shutdown ✅ → خروج

مدت زمان: معمولاً < 2 ثانیه
```

### حالت تایم‌اوت
```
اجرا → Hook طول می‌کشد → Timeout → خروج ⚠️

مدت زمان: = SHUTDOWN_TIMEOUT_SECONDS
```

### حالت اجباری
```
اجرا → سومین Ctrl+C → خروج فوری ❌

مدت زمان: فوری
```

---

## 📋 بهترین روش‌ها

### ✅ انجام دهید

1. **Hooks کوتاه نگاه دارید**
```go
// خوب
shutdownMgr.RegisterHook(func() error {
    file.Close()
    return nil
})
```

2. **Error handling داشته باشید**
```go
// خوب
if err := db.Close(); err != nil {
    return fmt.Errorf("failed to close db: %w", err)
}
```

3. **Resource cleanup انجام دهید**
```go
// خوب
defer file.Close()
defer db.Close()
```

### ❌ نکنید

1. **Infinite loops در hooks**
```go
// بد
shutdownMgr.RegisterHook(func() error {
    for {  // ← اینجا مسئله‌ای است
        doSomething()
    }
})
```

2. **Long-running operations**
```go
// بد
shutdownMgr.RegisterHook(func() error {
    time.Sleep(10 * time.Second)  // ← ممکن timeout شود
    return nil
})
```

3. **Panics در hooks**
```go
// بد
shutdownMgr.RegisterHook(func() error {
    panic("something went wrong")  // ← برنامه crash می‌شود
})
```

---

## 🐳 Docker Integration

### Docker Compose
```yaml
services:
  gold-analyzer:
    build: .
    restart: unless-stopped
    environment:
      - SHUTDOWN_TIMEOUT_SECONDS=10
    stop_grace_period: 15s  # Docker صبر می‌کند
```

### سیگنال‌های Docker
```bash
# Docker SIGTERM می‌فرستد
docker stop container_name

# برنامه graceful shutdown اجرا می‌کند
# داده‌ها ذخیره می‌شوند
# منابع بسته می‌شوند
```

---

## 🔍 Logging Shutdown

### Log Entry Example
```
[2025-12-14 22:53:14] SHUTDOWN: برنامه با موفقیت متوقف شد - آخرین سیگنال: HOLD
```

### مشاهدهٔ Log
```bash
# مشاهدهٔ زنده
tail -f signals.log

# جستجو برای shutdown
grep SHUTDOWN signals.log
```

---

## 🆘 مشکل‌گیری

### مشکل: برنامه بلافاصله خروج می‌کند

**دلیل**: Shutdown timeout خیلی کوتاه است

**راه حل**:
```bash
SHUTDOWN_TIMEOUT_SECONDS=30 ./analyzer
```

---

### مشکل: منابع بسته نشده‌اند

**دلیل**: Hook ثبت نشده است

**راه حل**:
```go
shutdownMgr.RegisterHook(func() error {
    return closeResources(cfg)
})
```

---

### مشکل: Timeout timeout می‌شود

**دلیل**: Hook‌ها خیلی طول می‌کشند

**راه حل**: Hooks را بهتر کنید یا timeout را زیاد کنید

---

## 📊 Performance

### Shutdown Speed
```
عادی:        0.5 - 2 ثانیه
با Hooks:     2 - 5 ثانیه
با Timeout:   = SHUTDOWN_TIMEOUT_SECONDS
```

### Resource Usage
```
Memory:      تقریباً صفر (تمام منابع بسته‌شده‌اند)
CPU:         صفر (برنامه متوقف‌شده است)
Files:       بسته‌شده‌اند
Logs:        ذخیره‌شده‌اند
```

---

## 🎓 مثال کامل

### راه‌اندازی
```go
package main

import (
    "gold-analyzer/shutdown"
)

func main() {
    // ایجاد manager
    mgr := shutdown.NewManager()
    
    // ثبت hooks
    mgr.RegisterHook(func() error {
        fmt.Println("بستن دیتابیس...")
        return nil
    })
    
    mgr.RegisterHook(func() error {
        fmt.Println("ذخیرهٔ داده‌ها...")
        return nil
    })
    
    // شروع
    mgr.Start()
    
    // منتظر سیگنال
    sig := mgr.WaitForShutdown()
    fmt.Printf("سیگنال: %v\n", sig)
    
    // Shutdown
    mgr.Shutdown(5 * time.Second)
}
```

---

## ✨ نکات نهایی

1. **Graceful shutdown اجباری است**
   - اگر Ctrl+C بزنید، منابع درست بسته می‌شوند
   - داده‌ها ضائع نمی‌شوند

2. **Hooks انعطاف‌پذیرند**
   - هر تعداد hook می‌توانید اضافه کنید
   - به ترتیب اجرا می‌شوند

3. **Timeout حفاظت می‌کند**
   - اگر hook‌ها طول بکشند، timeout می‌کند
   - برنامه نمی‌ماند "hanging"

4. **Docker-friendly**
   - SIGTERM را پردازش می‌کند
   - با orchestration اچ ظاهراً کار می‌کند

---

**نسخهٔ**: 1.0.0  
**تاریخ**: دسامبر 2025  
**وضعیت**: Production Ready ✅