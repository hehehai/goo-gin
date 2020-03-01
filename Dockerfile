FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/app
COPY . /go/app
RUN go mod download
RUN CGGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

CMD ["./go-gin-example"]

