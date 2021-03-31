FROM golang:1.15 as build

# Set the Current Working Directory inside the container
WORKDIR /build

COPY . .

# Build the Go app
RUN go build -o /app ./cmd/server

ENTRYPOINT ["/app"]

FROM gcr.io/distroless/base-debian10

COPY --from=build /app /

# This container exposes port 8080 to the outside world
EXPOSE 8888

# Run the binary program produced by `go install`
ENTRYPOINT ["/app"]