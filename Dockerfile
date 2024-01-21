# ステージ１
FROM golang:1.18.1-alpine3.15 AS go
WORKDIR /app
COPY app/* ./
RUN go mod download 
RUN go build -o main ./src/main.go

# ステージ２
FROM alpine:3.15
WORKDIR /app
COPY --from=go /app/main .
USER 1001
CMD [ "/app/main" ]