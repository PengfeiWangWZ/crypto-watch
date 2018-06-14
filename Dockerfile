FROM golang:1.10 as builder
WORKDIR /go/src/github.com/PengfeiWang/crypto-watch
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o crypto-watch

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/PengfeiWang/crypto-watch
COPY --from=builder /go/src/github.com/PengfeiWang/crypto-watch/crypto-watch .
CMD ["./crypto-watch"]