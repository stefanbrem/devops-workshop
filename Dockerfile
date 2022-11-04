# syntax=docker/dockerfile:1
FROM golang:1.16 as builder
WORKDIR /app/
COPY . .
# options are needed to run on alpine linux!
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-w -s" -o helloworld

FROM alpine:3.13 as final
RUN apk --no-cache add dumb-init
WORKDIR /app/
COPY --from=builder /app/ .
CMD ["/app/helloworld"]
