FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/src/app/out

EXPOSE 3000

CMD ["/usr/src/app/out"]
