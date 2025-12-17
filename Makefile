.PHONY: help build run test bench clean fmt lint install-tools

# متغیرها
BINARY_NAME=analyzer
GO=go
GOFLAGS=-v
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

help:
	@echo "🏆 Gold Analyzer - دستورات دستیار"
	@echo ""
	@echo "دستورات اصلی:"
	@echo "  make build          - کامپایل برنامه"
	@echo "  make run            - اجرای برنامه"
	@echo "  make test           - اجرای تست‌ها"
	@echo "  make bench          - اجرای بنچمارک‌ها"
	@echo "  make clean          - پاک کردن فایل‌های تولید شده"
	@echo "  make fmt            - فرمت کردن کد"
	@echo "  make lint           - بررسی کیفیت کد"
	@echo "  make install-tools  - نصب ابزارهای لازم"
	@echo "  make run-with-log   - اجرا با logging"
	@echo "  make run-with-charts - اجرا با تولید نمودارها"
	@echo "  make help           - نمایش این پیام"

build:
	@echo "🔨 درحال کامپایل..."
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) ./cmd/main.go
	@echo "✅ کامپایل کامل شد: ./$(BINARY_NAME)"

run: build
	@echo "🚀 درحال اجرا..."
	./$(BINARY_NAME)

run-with-log: build
	@echo "🚀 درحال اجرا با logging..."
	LOG_FILE="signals.log" ./$(BINARY_NAME)

run-with-shutdown-timeout: build
	@echo "🚀 درحال اجرا با timeout shutdown مخصوص..."
	SHUTDOWN_TIMEOUT_SECONDS=10 ./$(BINARY_NAME)

run-interactive: build
	@echo "🚀 درحال اجرا حالت تعاملی..."
	@echo "💡 برای متوقف کردن، Ctrl+C را فشار دهید..."
	./$(BINARY_NAME)

run-with-charts: build
	@echo "📊 درحال اجرا با تولید نمودارها..."
	CHART_OUTPUT_DIR="./charts" ./$(BINARY_NAME)

test:
	@echo "🧪 اجرای تست‌ها..."
	$(GO) test ./test -v
	@echo "✅ تست‌ها تکمیل شدند"

test-shutdown:
	@echo "🧪 اجرای تست‌های Shutdown..."
	$(GO) test ./test -v -run Shutdown
	@echo "✅ تست‌های Shutdown تکمیل شدند"

bench:
	@echo "📊 اجرای بنچمارک‌ها..."
	$(GO) test ./test -bench=. -benchmem
	@echo "✅ بنچمارک‌ها تکمیل شدند"

bench-cpu:
	@echo "📊 بنچمارک CPU..."
	$(GO) test ./test -bench=. -benchmem -cpuprofile=cpu.prof
	go tool pprof -http=:8080 cpu.prof

bench-mem:
	@echo "📊 بنچمارک حافظه..."
	$(GO) test ./test -bench=. -benchmem -memprofile=mem.prof
	go tool pprof -http=:8080 mem.prof

fmt:
	@echo "🎨 فرمت کردن کد..."
	$(GO) fmt ./...
	@echo "✅ کد فرمت شد"

lint:
	@echo "🔍 بررسی کیفیت کد..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "⚠️  golangci-lint نصب نشده است. دستور زیر رو اجرا کن:"; \
		echo "  make install-tools"; \
	fi

vet:
	@echo "🔎 Go vet بررسی..."
	$(GO) vet ./...
	@echo "✅ vet بررسی تکمیل شد"

install-tools:
	@echo "⚙️  نصب ابزارهای توسعه..."
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ ابزارها نصب شدند"

clean:
	@echo "🧹 پاک کردن فایل‌های تولید شده..."
	rm -f $(BINARY_NAME)
	rm -f signals.log
	rm -rf ./charts
	rm -f *.prof
	rm -f *.test
	$(GO) clean
	@echo "✅ پاک شد"

deps:
	@echo "📦 دانلود وابستگی‌ها..."
	$(GO) mod download
	$(GO) mod tidy
	@echo "✅ وابستگی‌ها دانلود شدند"

mod-graph:
	@echo "📊 نمودار وابستگی‌ها:"
	$(GO) mod graph

check: fmt vet test lint
	@echo "✅ تمام بررسی‌ها تکمیل شدند"

check-all: clean fmt vet test bench lint
	@echo "✅ تمام بررسی‌های جامع تکمیل شدند"

dev: clean build run-with-log

dev-charts: clean build run-with-charts

dev-interactive: clean build run-interactive

release: clean test
	@echo "📦 ساخت نسخه release..."
	GOOS=linux GOARCH=amd64 $(GO) build -o analyzer-linux-amd64 ./cmd/main.go
	GOOS=darwin GOARCH=amd64 $(GO) build -o analyzer-darwin-amd64 ./cmd/main.go
	GOOS=darwin GOARCH=arm64 $(GO) build -o analyzer-darwin-arm64 ./cmd/main.go
	GOOS=windows GOARCH=amd64 $(GO) build -o analyzer-windows-amd64.exe ./cmd/main.go
	@echo "✅ نسخه release ساخته شدند"

info:
	@echo "📋 اطلاعات پروژه:"
	@echo "  Go Version: $(shell $(GO) version)"
	@echo "  OS: $(shell uname -s)"
	@echo "  Arch: $(shell uname -m)"
	@echo "  Binary: $(BINARY_NAME)"
	@echo "  Test Files: $(shell find . -name '*_test.go' | wc -l)"
	@echo "  Go Files: $(shell find . -name '*.go' | wc -l)"

all: clean fmt lint test build
	@echo "✅ تمام مراحل تکمیل شدند"

# Shutdown management
shutdown-help:
	@echo "🛑 راهنمای Graceful Shutdown:"
	@echo ""
	@echo "دستورات Shutdown:"
	@echo "  make run-interactive         - اجرا و Ctrl+C برای shutdown"
	@echo "  make run-with-shutdown-timeout - اجرا با timeout مخصوص"
	@echo "  make test-shutdown           - تست کردن shutdown mechanism"
	@echo ""
	@echo "متغیرهای Environment:"
	@echo "  SHUTDOWN_TIMEOUT_SECONDS=N   - تنظیم timeout (پیش‌فرض: 5)"
	@echo ""
	@echo "مثال:"
	@echo "  SHUTDOWN_TIMEOUT_SECONDS=15 ./analyzer"
