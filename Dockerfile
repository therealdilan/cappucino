FROM golang:1.16-alpine

WORKDIR /cappucino


COPY go.mod .
COPY main.go .
COPY app/index.html app/index.html
COPY app/static app/static

RUN go build -o bin . 

EXPOSE 6969

ENTRYPOINT [ "/cappucino/bin" ]

