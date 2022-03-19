# Start by selecting the base image for our service
FROM golang:1.16-alpine

# Creating the `app` directory in which the app will run 
RUN mkdir /app

# Move everything from root to the newly created app directory
ADD . /app

# Specifying app as our work directory in which
# futher instructions should run into
WORKDIR /app

# Download all neededed project dependencies
RUN go mod download

# Build the project executable binary
RUN go build -o main .

EXPOSE 2222

# Run/Starts the app executable binary
CMD ["/app/main"]