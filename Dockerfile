FROM golang:1.8.3-alpine3.6

RUN apk --update add git && rm -rf /var/cache/apk/*
RUN go get -u github.com/labstack/echo/...

COPY . .
EXPOSE 8000

CMD ["go","run","main.go"]
