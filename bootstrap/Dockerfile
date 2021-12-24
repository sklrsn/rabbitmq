FROM golang:latest

WORKDIR /go/src/github.com/sklsn/bootstrap

ADD . /go/src/github.com/sklsn/bootstrap/

RUN go mod init github.com/sklsn/bootstrap && \
    go mod download && \
    go mod tidy

RUN go build -o /opt/rabbitmq/bootstrap

RUN chmod +x /opt/rabbitmq/bootstrap

RUN useradd executor 

CMD [ "/opt/rabbitmq/bootstrap" ]