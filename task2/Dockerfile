FROM golang:1.20
RUN mkdir /app
WORKDIR /app
COPY /main.go /app
RUN GO111MODULE=off go build -o main .
CMD ["/app/main"]
