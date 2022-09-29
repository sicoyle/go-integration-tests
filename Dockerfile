ARG GO_IMAGE=golang:1.19.1-alpine3.16

# builder stage
FROM ${GO_IMAGE} AS builder
ENV CGO_ENABLED=0
ADD . /service/
WORKDIR /service
RUN go mod download
RUN go build main.go

# final stage
FROM scratch
COPY --from=builder /service/main main
ENTRYPOINT ["/main"]