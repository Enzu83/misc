# Simple Alpine Docker Container

This is a basic Docker container based on Alpine Linux that demonstrates file copying and exploration.

## Building the Container

```bash
cd docker/simple-alpine
docker build -t simple-alpine .
```

## Running the Container

### Interactive Mode (Recommended for file exploration)
```bash
docker run -it simple-alpine
```

This will start the container with an interactive shell where you can explore the filesystem.

## Exploring Files in the Container

Once you're inside the container (after running `docker run -it simple-alpine`), you can use these commands to explore the files:

### Basic File Listing
```bash
# List files in current directory
ls -la

# List files in root directory
ls -la /

# List files recursively with tree structure
tree /app
```

### File Content Viewing
```bash
# View file content
cat sample.txt

# View file content with pagination
less sample.txt

# View first few lines
head sample.txt

# View last few lines
tail sample.txt
```

### File System Navigation
```bash
# Navigate to different directories
cd /
cd /app
cd /etc

# Show current directory
pwd

# Find files by name
find / -name "sample.txt" 2>/dev/null

# Find files by type
find /app -type f  # files only
find /app -type d  # directories only
```

### File Information
```bash
# Show file details
ls -la sample.txt

# Show file type
file sample.txt

# Show disk usage
du -sh /app
```

### System Information
```bash
# Show Alpine version
cat /etc/alpine-release

# Show system information
uname -a

# Show mounted filesystems
df -h

# Show processes
ps aux
```

### Text Editing (if needed)
```bash
# Edit files with nano
nano sample.txt

# Or use vi (if available)
vi sample.txt
```

## Alternative: Exploring Without Interactive Mode

If you want to explore files without entering the container interactively:

```bash
# Execute single commands
docker run --rm simple-alpine ls -la /app
docker run --rm simple-alpine cat /app/sample.txt
docker run --rm simple-alpine tree /app

# Copy files from container to host
docker create --name temp simple-alpine
docker cp temp:/app/sample.txt ./copied-sample.txt
docker rm temp
```

## Exploring Container Layers

You can also explore the container's layers and metadata:

```bash
# Inspect the image
docker inspect simple-alpine

# View image history
docker history simple-alpine

# View image layers
docker image inspect simple-alpine --format='{{json .RootFS.Layers}}'
```
