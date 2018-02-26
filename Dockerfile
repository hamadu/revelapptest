FROM golang:1.10.0 

ENV GOPATH $GOPATH:/go/src 

RUN apt-get update && \
    apt-get upgrade -y

RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel && \
    go get github.com/coopernurse/gorp && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql

RUN mkdir /go/src/app

EXPOSE 9000
