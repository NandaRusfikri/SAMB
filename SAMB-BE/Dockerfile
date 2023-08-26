# Create a minimal Docker image with only the binary
FROM alpine:latest

# Update package repository and install required packages
RUN apk update && \
    apk add --no-cache bash tzdata && \
    rm -rf /var/cache/apk/*

# Set the timezone
ENV TZ=Asia/Jakarta

# Copy the binary from the previous build stage
COPY vdi-meter vdi-meter
COPY assets assets

# Set the entry point to the binary
ENTRYPOINT ["/vdi-meter"]

# Set the default command for the container
CMD ["./vdi-meter"]
