FROM golang:alpine as builder

# Install git.
# Git is required for fetching the dependencsies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/web

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .      
COPY --from=builder /app/static/ ./static
COPY --from=builder /app/templates/ ./templates  

# Expose port 8080 to the outside world
EXPOSE 8080

#Run the executable
CMD ["./main"]