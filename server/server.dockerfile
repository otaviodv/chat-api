FROM golang:latest as builder

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o serverApp .

RUN chmod +x /app/serverApp

FROM alpine:latest
RUN mkdir /app

COPY --from=builder /app/serverApp /app

# EXPOSE 80

CMD [ "/app/serverApp" ]