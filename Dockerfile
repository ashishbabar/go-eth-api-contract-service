FROM golang:1.16-alpine
WORKDIR $GOPATH/src/github.com/ashishbabar/go-eth-api-contract-service

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 3000

CMD [ "go-eth-api-contract-service" ]