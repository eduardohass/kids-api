FROM golang:1.23rc2-bullseye AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
    git \
    ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go clean -modcache
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-w -s" \
    -o kids-api ./cmd/api

FROM gcr.io/distroless/static-debian11:latest
COPY --from=builder /app/kids-api /
COPY --from=builder /app/internal/migrations /migrations

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/kids-api"]