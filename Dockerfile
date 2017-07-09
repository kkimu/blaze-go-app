FROM golang:1.8.3-alpine3.6

RUN apk --update add git
RUN apk --update add build-base
RUN apk --update add pkgconfig
RUN apk --update add ffmpeg-dev
RUN apk --update add ffmpeg && rm -rf /var/cache/apk/*

RUN go get -u github.com/opennota/screengen
RUN go get -u github.com/labstack/echo/...
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/k0kubun/pp


WORKDIR /go/src/github.com/kkimu/blaze-go-app
EXPOSE 8000

CMD ["go","run","main.go"]
