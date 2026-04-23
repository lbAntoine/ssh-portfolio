FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ssh-portfolio ./cmd/server


FROM gcr.io/distroless/static-debian12

COPY --from=builder /build/ssh-portfolio /ssh-portfolio

EXPOSE 2222

ENTRYPOINT ["/ssh-portfolio"]
