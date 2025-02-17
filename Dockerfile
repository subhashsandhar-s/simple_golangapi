FROM golang:1.24.0

WORKDIR /app
COPY . /app
RUN go mod tidy
CMD ["go", "run", "main.go"]