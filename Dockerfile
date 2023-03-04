FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/backend ./cmd/server/main.go

RUN apk add upx
RUN upx --ultra-brute /app/backend

FROM scratch AS prod
COPY --from=builder /app/backend /bin/backend

EXPOSE 8080

ENTRYPOINT ["/bin/backend"]
