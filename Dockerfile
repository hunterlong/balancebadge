FROM golang

RUN apk update && apk add git

COPY . $GOPATH/src/github.com/hunterlong/balancebadge/
WORKDIR $GOPATH/src/github.com/hunterlong/balancebadge/
RUN go get -d -v
RUN go install
WORKDIR /app
VOLUME /app

EXPOSE 9090

ENTRYPOINT balancebadge