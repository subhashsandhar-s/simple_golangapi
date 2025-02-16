FROM golang:1.24.0

WORKDIR /app
COPY . /app
RUN go get "github.com/gorilla/mux"
CMD ["go", "run", "main.go"]