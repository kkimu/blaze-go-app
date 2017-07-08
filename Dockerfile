FROM golang:1.8.3-alpine3.6

RUN apk --update add git && rm -rf /var/cache/apk/*
RUN go get -u github.com/labstack/echo/...
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm

WORKDIR /go/src/github.com/kkimu/blaze-go-app
EXPOSE 8000

CMD ["go","run","main.go"]
