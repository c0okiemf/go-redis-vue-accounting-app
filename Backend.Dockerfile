FROM golang:1.20

WORKDIR /app

COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod download

COPY ./backend .

EXPOSE 8080

CMD ["go", "run", "main.go"]
