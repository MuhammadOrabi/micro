FROM golang

RUN go get golang.org/x/tools/cmd/goimports
RUN go get database/sql
RUN go get -u github.com/goadesign/goa/...
RUN go get github.com/joho/godotenv
RUN go get github.com/lib/pq
RUN go get github.com/streadway/amqp

WORKDIR /go/src/micro

COPY . .

EXPOSE 3000

CMD go build && ./micro