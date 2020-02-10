## Build with: docker build -t dataregistration .
## Run with: docker run -i -p 8080:8080 -p 8088:8088 dataregistration

# FIRST STAGE:  build the app.
FROM golang:1.13 AS build-app
WORKDIR /build/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# Dependencies are downloaded only when go.mod or go.sum changes.
RUN go mod download

# Copy the rest of the source files.
COPY . .
RUN go test ./internal/app
RUN go build -o dynamic-geoassets-api ./cmd/dynamic-geoassets-api

# SECOND STAGE: create the app runtime image.
FROM ubuntu:bionic

COPY --from=build-app /build/app/dynamic-geoassets-api /app/
COPY --from=build-app /build/app/static /app/static
COPY --from=build-app /build/app/templates /app/templates

WORKDIR /app
USER nobody:nogroup

ENTRYPOINT ["/app/dynamic-geoassets-api"]
