FROM golang:1.21-alpine3.18 as build-env

# Install git
RUN apk update apk install
RUN apk add --no-cache git

# Create a non-root user
RUN adduser -D -g '' appuser

# Set the working directory
RUN mkdir /app
WORKDIR /app

# Copy the rest of the application source code
COPY ./webserver/go.mod .
COPY ./webserver/go.sum .
COPY ./webserver/main.go .

# Build the application
RUN apk --no-cache add ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/app ./main.go

# Use a minimalistic base image
FROM alpine:3.18

# Create a non-root user
RUN adduser -D -g '' appuser

# Copy necessary files from the build environment including the config directory
WORKDIR /
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env --chown=appuser /app/app /webserver/app
COPY --chown=appuser ./website /website

# Build the website
WORKDIR /website
RUN rm -rf node_modules
RUN rm -rf package-lock.json
RUN apk add --no-cache nodejs npm
RUN npm install
WORKDIR /

# Permissions to expose port 80
RUN apk --no-cache add libcap && setcap 'cap_net_bind_service=+ep' /webserver/app

# Add the appuser to the passwd file
RUN echo "appuser:x:10001:10001::/:" >> /etc/passwd

# Switch to the non-root user
USER appuser

# Set the entrypoint
ENTRYPOINT ["/webserver/app"]
