FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
# EXPOSE 8080
RUN chmod +x /app/main
ENTRYPOINT ["./main"]
