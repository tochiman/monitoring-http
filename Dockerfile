FROM golang:1.19-alpine as go
ENV CGO_ENABLED 0
WORKDIR /go/src
RUN apk update && apk add git
COPY ./ ./
RUN go mod download && go build -o /go/bin/main ./main.go

FROM gcr.io/distroless/static-debian11
COPY --from=go /go/bin/main /
CMD ["/main"]