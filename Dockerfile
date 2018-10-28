
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY ./data /data
ENTRYPOINT ./app
LABEL Name=go-geolocator Version=0.0.1
EXPOSE 1323

ARG city_mmdb_path=data/GeoLite2-City.mmdb
ENV CITY_MMDB_PATH=$city_mmdb_path
ENV PORT=$PORT

CMD [ "server" ]
