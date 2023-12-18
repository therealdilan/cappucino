FROM golang:1.16-alpine

WORKDIR /app


COPY go.mod .
COPY main.go .
COPY app/index.html .
COPY app/static ./static

RUN go build -o bin . 

EXPOSE 6969

ENTRYPOINT [ "/cappucino/bin" ]

