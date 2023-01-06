FROM golang:1.18.9-alpine3.17
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get 
RUN go build -o main .
CMD ["/app/main"]
