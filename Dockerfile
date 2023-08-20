# Use an official Go runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code into the container at the working directory
COPY . /app

# Specify the command to run when the container starts
CMD ["go", "run", "main.go"]
