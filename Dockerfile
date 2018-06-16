FROM golang

COPY . $GOPATH/src/github.com/hunterlong/balancebadge/
WORKDIR $GOPATH/src/github.com/hunterlong/balancebadge
RUN go get
RUN go install

EXPOSE 9090

ENTRYPOINT balancebadge