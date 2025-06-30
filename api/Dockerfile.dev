FROM golang:alpine
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.* .
RUN go mod download
CMD ["air", "-c", ".air.toml"]
