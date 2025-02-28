FROM golang:latest as builder

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o botApp .

RUN chmod +x /app/botApp

FROM alpine:latest
RUN mkdir /app

COPY --from=builder /app/botApp /app

CMD [ "/app/botApp" ]