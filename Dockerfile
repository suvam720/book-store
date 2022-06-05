FROM  golang:1.18-alpine3.15
WORKDIR /myapp
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app
EXPOSE 8080
CMD ["./app"]
