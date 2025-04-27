FROM golang:1.24-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/app/main.go

FROM alpine AS runner

WORKDIR /app

COPY --from=build /build/main ./main
COPY --from=build build/cmd/app/config.yaml ./config.yaml

RUN chmod +x ./main

CMD ["/app/main", "-config", "config.yaml"]