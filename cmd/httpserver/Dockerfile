# Make sure to specify the same Go version as the one in the go.mod file.
# For example, golang:1.22.1-alpine.
FROM golang:1.24-alpine

WORKDIR /app


COPY . .

RUN go build -o svr cmd/httpserver/main.go

EXPOSE 8080
CMD [ "./svr" ]