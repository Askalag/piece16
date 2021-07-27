FROM golang:alpine

WORKDIR /go/src/
COPY ./ ./
RUN go build -o piece16 cmd/main.go
EXPOSE 8088
CMD ["./piece16"]