FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/app
COPY . /go/app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

EXPOSE 9090
ENTRYPOINT ["./app"]

