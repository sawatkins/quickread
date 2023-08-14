FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o upfast .

EXPOSE 8080

CMD ["/app/upfast", "-prefork=false", "-port=:8080"]

