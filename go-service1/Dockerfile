
# Select base image as golang versio 1.11.2
FROM golang:1.11.2

# Set working directory for entry point to path
WORKDIR /go/src/go-service1

# Add project directory to path
ADD . /go/src/go-service1

# Execute command for get library and run command for init dependency
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init -v

# Set default run command after container created and started
CMD ["go","run","main.go"]
