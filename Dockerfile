FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o quickread .

EXPOSE 8080

CMD ["/app/quickread", "-prefork=false", "-port=:8080"]

