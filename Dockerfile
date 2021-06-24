FROM golang:1.11

WORKDIR /go/src/app
COPY . .

ENV PATH $GOPATH/bin:$PATH

RUN go get github.com/go-task/task
RUN go install github.com/go-task/task
RUN task test
RUN task cover-lexer
RUN bash <(curl -s https://codecov.io/bash) -t a500981d-b8e2-4ea8-ae17-b1d4065e9015
