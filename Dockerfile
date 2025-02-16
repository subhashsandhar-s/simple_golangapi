FROM golang:1.23.4

WORKDIR /app
COPY . /app
RUN go get "github.com/gorilla/mux"
CMD ["go", "run", "main.go"]