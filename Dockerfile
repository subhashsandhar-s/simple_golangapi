FROM golang:1.24.0

WORKDIR /app
COPY . /app
CMD ["go", "run", "main.go"]