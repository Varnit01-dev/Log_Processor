FROM golang:alpine

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

COPY log_collector.go /app/
COPY logparser.go /app/

RUN go build -o log-processor log_collector.go logparser.go

CMD ["log-processor"]
