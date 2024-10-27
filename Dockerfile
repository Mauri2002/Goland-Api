FROM golang:1.23.2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *go ./

RUN go build 

EXPOSE 5050

CMD ["/app/RestApiPract1"]
