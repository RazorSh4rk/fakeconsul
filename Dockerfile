FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
ENV PORT=8500
EXPOSE 8500

CMD ["./main"]
