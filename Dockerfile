FROM golang:1.18-alpine as build-env
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o cloudsignfakeserver
FROM alpine:latest
COPY --from=build-env /app/cloudsignfakeserver .
EXPOSE 18080/tcp
USER 1001
ENTRYPOINT ["./cloudsignfakeserver"]
