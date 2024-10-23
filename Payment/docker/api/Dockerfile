FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY src/api/go.* ./
RUN go mod download

COPY src/api ./

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]