# -----------------------------
# Stage 1: Build the Go application
# -----------------------------
FROM golang:1.21.3-alpine AS builder


WORKDIR $GOPATH/src

# Copy all files from the current directory to the working directory in the container
COPY . .

# Bypass the default Go proxy and fetch directly from the source
ENV GOPROXY=direct

# Install 'git' and other dependencies necessary for building the application
RUN apk add --no-cache git 

# Fetch and tidy up the Go module dependencies
RUN go mod tidy -x

# Compile the Go application with specific settings
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/grender main.go

# -----------------------------
# Stage 2: Create the final image
# -----------------------------
FROM scratch

# Copy the compiled binary from the builder stage to the final container
COPY --from=builder /go/bin/grender /go/bin/grender

# Expose port 8080 for the application
EXPOSE 8080

# Set the command to run when the container starts
ENTRYPOINT [ "/go/bin/grender" ]