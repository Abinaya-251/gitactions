ROM golang:1.16.4-buster AS builder
  
WORKDIR /app
COPY main.go .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go

FROM debian:buster-slim
COPY --from=builder /app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]
