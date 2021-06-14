FROM golang:1.15-alpine as builder

LABEL org.opencontainers.image.source = "https://github.com/blyndusk/image-resizer"

# ----- SETUP -----

# Enable Go modules
ENV GO111MODULE=on

# Set the current working with go absolute path
WORKDIR /go/src/github.com/blyndusk/image-resizer/api

# ----- INSTALL -----

# Copy go.mod + go.sum for install before full copy
COPY api/go.mod .

COPY api/go.sum .

# Download all dependencies
RUN go mod download -x

# ----- COPY + RUN -----

# Copy the source from the current directory to the container
COPY api/ .

# Build app into specific folder
RUN go build -o ./tmp/main ./

# ----- ALPINE -----

FROM alpine

# Copy binary
COPY --from=builder /go/src/github.com/blyndusk/image-resizer/api/ /image-resizer/api/

# Set current directory
WORKDIR /image-resizer/api/

# Use executable
ENTRYPOINT [ "/image-resizer/api/tmp/main" ]
