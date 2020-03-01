FROM golang:1.14.0 AS gobuild

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go/cache

COPY go.mod .
COPY go.sum .

RUN go mod download

WORKDIR /go/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch AS prod

COPY --from=gobuild /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=gobuild /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=gobuild /go/app/app /
COPY --from=gobuild /go/app/conf/app.ini /conf/

EXPOSE 9090
ENTRYPOINT ["/app"]

