FROM golang:1.22.4-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=build /app/myapp .

ENTRYPOINT ["/app/myapp"]