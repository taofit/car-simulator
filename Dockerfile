FROM golang:1.19.4 
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-car
# Expose port
EXPOSE 4000
# Start app
CMD ["/docker-car"]