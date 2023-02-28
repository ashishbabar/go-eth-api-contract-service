FROM golang:1.18-alpine
WORKDIR $GOPATH/src/github.com/ashishbabar/go-eth-api-contract-service

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN go get -tags musl -d -v ./...

RUN go build -tags musl -o app ./cmd/http.go

EXPOSE 3000

CMD [ "./app" ]