# Changelog

تمام تغییرات مهم در این پروژه در این فایل ثبت می‌شوند.

## [1.0.0] - 2025-12-14

### ✨ ویژگی‌های اضافه شده
- 🚀 سیستم نظارت خودکار هر یک دقیقه
- 📊 محاسبهٔ اندیکاتورهای تکنیکال (RSI، MACD، ATR)
- 🎯 تولید سیگنال‌های معاملاتی (BUY/SELL/HOLD)
- 💾 سیستم logging برای ذخیرهٔ سیگنال‌ها
- ⚙️ تنظیمات قابل سفارشی‌سازی
- 🔄 Retry logic برای handling rate limiting
- 📝 مستندات جامع (README، QUICK_START)
- 🧪 Unit tests و Benchmark tests
- 🐳 Docker و docker-compose support
- 🛠️ Makefile با 20+ دستور

### 🔧 اصلاحات
- ✅ رفع خطای کامپایل (اسم ماژول در go.mod)
- ✅ بهبود error handling
- ✅ اضافه کردن User-Agent برای درخواست‌ها
- ✅ exponential backoff برای retry logic

### 📚 مستندات
- ✅ README.md - مستندات کامل
- ✅ QUICK_START.md - شروع سریع
- ✅ PROJECT_STATUS.md - وضعیت پروژه
- ✅ .env.example - نمونهٔ تنظیمات

### 🏗️ معماری
- ✅ Modular structure
- ✅ Clean code principles
- ✅ Separation of concerns
- ✅ Configuration management

### 🚀 Deployment
- ✅ Dockerfile برای containerization
- ✅ docker-compose.yml برای orchestration
- ✅ Multi-platform build support

## [0.1.0] - Unreleased (اولین مفهوم)

### شروع
- اولین نسخهٔ پروژه
- اندیکاتورهای پایه‌ای
- استراتژی سادهٔ معاملاتی

---

## نحوهٔ استفاده از این Changelog

این Changelog بر اساس [Keep a Changelog](https://keepachangelog.com/) تنظیم شده است.

### نسخه‌گذاری
ما از [Semantic Versioning](https://semver.org/) استفاده می‌کنیم:
- MAJOR: تغییرات ناسازگار
- MINOR: ویژگی‌های جدید (سازگار)
- PATCH: اصلاحات و بهبودها

### برچسب‌ها
- 🚀 New features
- 🔧 Changes
- ✅ Fixes
- 📚 Documentation
- ⚠️ Deprecated
- 🔒 Security
- 🐛 Bug fixes

---

## نقشهٔ راه (Roadmap)

### فاز بعدی (v1.1.0)
- [ ] Telegram notifications
- [ ] Advanced indicators
- [ ] Backtesting engine
- [ ] Performance improvements

### فاز دوم (v1.2.0)
- [ ] Web Dashboard
- [ ] Database support
- [ ] Multi-pair analysis
- [ ] Risk management

### فاز سوم (v2.0.0)
- [ ] Machine Learning integration
- [ ] Mobile app
- [ ] WebSocket support
- [ ] Advanced charting

---

## Contributing

برای کمک به این پروژه:
1. یک Issue باز کنید
2. یک Fork بسازید
3. تغییرات را commit کنید
4. Pull Request بسازید

هنگام push، لطفاً CHANGELOG را به‌روز کنید!

---

## License

MIT License - برای اطلاعات بیشتر LICENSE فایل را ببینید.

---

**نسخهٔ فعلی**: 1.0.0  
**آخرین بروزرسانی**: دسامبر 14، 2025