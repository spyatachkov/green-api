FROM golang:1.25-alpine AS builder

WORKDIR /opt/app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o greenapi-backend ./cmd/main.go

FROM alpine:latest

WORKDIR /opt/app

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

COPY --from=builder /opt/app/greenapi-backend .

RUN chown -R appuser:appuser /opt/app

USER appuser

EXPOSE 8080

CMD ["./greenapi-backend"]
