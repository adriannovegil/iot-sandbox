# https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles/
# Build
FROM golang:latest as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
# Build application
RUN make clean; make build

# Run
FROM scratch
COPY --from=builder /build/load-agent /app/
WORKDIR /app
CMD ["./load-agent"]
