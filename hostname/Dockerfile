FROM golang:1.5
COPY main.go $GOPATH/src/nds44/hostname/
RUN go get nds44/hostname
RUN go install nds44/hostname
CMD ["hostname"]
