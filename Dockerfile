FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o goipbot ./cmd/main.go

CMD ["./goipbot"]
