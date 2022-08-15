FROM golang:1.19-alpine

ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]