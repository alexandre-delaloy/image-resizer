FROM golang:1.15-alpine as builder

LABEL org.opencontainers.image.source = "https://github.com/blyndusk/image-resizer"

# ----- SETUP -----

# Enable Go modules
ENV GO111MODULE=on

# Set the current working with go absolute path
WORKDIR /go/src/github.com/blyndusk/image-resizer/worker

# ----- INSTALL -----

# Copy go.mod + go.sum for install before full copy
COPY worker/go.mod .

COPY worker/go.sum .

# Download all dependencies
RUN go mod download -x

# ----- COPY + RUN -----

# Copy the source from the current directory to the container
COPY worker/ .

# Build app into specific folder
RUN go run main.go