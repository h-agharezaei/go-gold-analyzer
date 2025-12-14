# Build stage
FROM golang:1.25.2-alpine AS builder

WORKDIR /app

# نصب وابستگی‌ها
COPY go.mod ./
RUN go mod download

# کپی کردن کد
COPY . .

# کامپایل
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o analyzer ./cmd/main.go

# Runtime stage
FROM alpine:latest

# نصب ssl certificates برای HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# کپی کردن binary از builder
COPY --from=builder /app/analyzer .

# نصب tzdata برای timezone support
RUN apk add --no-cache tzdata

# Environment variables
ENV SYMBOL=GC=F
ENV INTERVAL=1h
ENV RANGE=7d
ENV CHECK_INTERVAL_MINUTES=1
ENV RSI_PERIOD=14
ENV ATR_PERIOD=14
ENV RSI_BUY_LOWER=40
ENV RSI_BUY_UPPER=55
ENV RSI_SELL_THRESHOLD=65
ENV LOG_FILE=/var/log/gold-analyzer/signals.log
ENV ENABLE_NOTIFICATIONS=0

# ایجاد دایرکتوری برای log
RUN mkdir -p /var/log/gold-analyzer

# اجرا
CMD ["./analyzer"]
