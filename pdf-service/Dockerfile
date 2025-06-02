# Stage 1: Base - Install wkhtmltopdf
FROM ubuntu:20.04 AS base
# Update package list and upgrade packages using an alternative mirror
RUN sed -i 's|http://archive.ubuntu.com/ubuntu/|http://mirrors.kernel.org/ubuntu/|g' /etc/apt/sources.list \
    && apt-get update && apt-get -y upgrade
# Install wget
RUN apt-get install -y wget
# Download wkhtmltopdf package for Ubuntu
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.focal_amd64.deb
# Install wkhtmltopdf dependencies
RUN apt-get install -y \
    fontconfig \
    libfreetype6 \
    libjpeg8 \
    libpng16-16 \
    libssl1.1 \
    libx11-6 \
    libxext6 \
    libxcb1 \
    libxrender1 \
    xfonts-75dpi \
    xfonts-base
# Install wkhtmltopdf
RUN dpkg -i wkhtmltox_0.12.6-1.focal_amd64.deb
# Verify wkhtmltopdf installation
RUN wkhtmltopdf --version
# Stage 2: Build - Build the Go application
FROM golang:1.21.6-alpine AS build
# Switches to /app as the working directory, similar to 'cd'
WORKDIR /app
# If you have a go.mod and go.sum file in your project, uncomment lines 13, 14, 15
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# Builds the current project to a binary file called motor-insurance
RUN GOOS=linux go build -o ./out/pdf-service .
# Stage 3: Final - Create the final image with wkhtmltopdf and the Go application
FROM ubuntu:20.04 AS final
# Install required dependencies including libssl1.1 for wkhtmltopdf
RUN sed -i 's|http://archive.ubuntu.com/ubuntu/|http://mirrors.kernel.org/ubuntu/|g' /etc/apt/sources.list \
    && apt-get update && apt-get install -y \
    fontconfig \
    libfreetype6 \
    libjpeg8 \
    libpng16-16 \
    libx11-6 \
    libxcb1 \
    libxext6 \
    libxrender1 \
    xfonts-75dpi \
    xfonts-base \
    libssl1.1 \
    ca-certificates \
    tzdata \
    --no-install-recommends && rm -rf /var/lib/apt/lists/*
# Copy wkhtmltopdf from the base stage
COPY --from=base /usr/local/bin/wkhtmltopdf /usr/local/bin/wkhtmltopdf
# Copy the Go application from the build stage
COPY --from=build /app/out/motor-insurance-docker /app/motor-insurance-docker
# Switches working directory to /app
WORKDIR "/app"
# Copy necessary files

ENV TZ=Asia/Kolkata
# Exposes the 8586 port from the container
EXPOSE 8000
# Runs the binary once the container starts
CMD ["./pdf-service"]