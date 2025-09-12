# Use Alpine Linux as the base image
FROM alpine:3.19

# Set working directory
WORKDIR /tst

# Copy a sample file to the container
COPY embedded/lib /tst/embedded/lib/

# Install basic utilities for file exploration
RUN apk add --no-cache \
    tree \
    less \
    nano

# Set default command to keep container running
CMD ["sh"]
