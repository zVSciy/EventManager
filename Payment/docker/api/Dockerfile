# ------------------------------------------------------ #

FROM golang:1.23-alpine AS build

RUN apk add --no-cache tzdata

WORKDIR /app

COPY src/api/go.* ./
RUN go mod download

COPY src/api ./
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/api/main.go

# ------------------------------------------------------ #

FROM build AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]

# ------------------------------------------------------ #

FROM build AS tests

RUN go test ./...

# ------------------------------------------------------ #

FROM scratch AS production

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /app/bin/main /

EXPOSE 3000

ENTRYPOINT [ "/main" ]