FROM golang:1.14.0 AS gobuild

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go/cache

COPY go.mod .
COPY go.sum .

RUN go mod download

WORKDIR /go/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 -ldflags="-s -w" -installsuffix go build -o app main.go

FROM scratch AS prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/app/app /
COPY --from=build /go/app/conf/app.ini /conf/

EXPOSE 9090
ENTRYPOINT ["/app"]

