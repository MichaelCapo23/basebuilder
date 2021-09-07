# Base Builder
FROM golang:1.15.2-alpine3.12 AS base-builder
WORKDIR /jwtserver
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd cmd/
COPY pkg pkg/
COPY internal internal/
RUN go test -short ./...
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Base Server
FROM scratch as base-server
COPY --from=base-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base-builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip

# api service
FROM base-builder as api-builder
RUN go build -o /go/bin/api cmd/main.go

FROM base-server as api
COPY --from=api-builder /go/bin/api api
ENTRYPOINT ["/api"]

# migration 
FROM base-builder as migration-builder
RUN mkdir /migration & \
    go build -o /migration/migration migration/main.go & \
    wait

FROM base-server as migration
COPY --from=migration-builder /migration /migration
ADD /migration/migrations /jwtserver/migration/migrations
WORKDIR /migration
ENTRYPOINT ["/migration/migration up"]