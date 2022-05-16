FROM  golang:1.18-alpine3.15
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN go mod download
RUN go build -o /app

CMD ["/app"]