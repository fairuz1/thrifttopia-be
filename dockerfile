FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .
RUN go build -v -o /usr/local/bin/app ./

EXPOSE 9990

CMD ["/usr/local/bin/app"]