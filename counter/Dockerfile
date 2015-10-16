FROM golang:1.5
COPY main.go $GOPATH/src/nds44/counter/
RUN go get nds44/counter
RUN go install nds44/counter
CMD ["counter"]
