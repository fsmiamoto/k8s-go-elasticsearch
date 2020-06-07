FROM golang:alpine AS builder
CMD mkdir /build
WORKDIR /build
COPY . .
ENV GO111MODULE=on
RUN go build -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /build/app /app/
ENV ELASTIC_APM_SERVER_URL=http://apm-server:8200
CMD ["./app"]
