

#FROM golang:1.17-alpine as build1
#WORKDIR /app1
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux go build -o rf-car.exe

FROM alpine:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /
#COPY --from=build1 /app1/app1.exe .
COPY tusd/tusd .

## if using http
#CMD ./tusd --hooks-http http://localhost:8081/write

## if using grpc
CMD ./tusd --hooks-grpc '127.0.0.1:8083' --hooks-grpc-retry 5 --hooks-grpc-backoff 2

