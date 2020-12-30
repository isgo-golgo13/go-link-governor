FROM golang:1.15 AS builder
ADD ./main.go main.go
ADD ./service service
# Fetch dependencies
RUN go get -d -v

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
MAINTAINER Tom Wallick <twallick@vivsoft.io>
COPY --from=builder /link_guardian_service /app/link_guardian_service
EXPOSE 7070
ENTRYPOINT ["/app/link_guardian_service"]