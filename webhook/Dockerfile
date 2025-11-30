

FROM golang:tip-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o webhook main.go


FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/webhook .
EXPOSE 9001
CMD ["./webhook"]


