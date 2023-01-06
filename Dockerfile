# Specify the base image we need for out go app
FROM golang:1.18.9-alpine3.17
# Create a directory for copy our project
RUN mkdir /app
# Copy our project to the docker directory
ADD . /app
# change /app as working directory 
WORKDIR /app
# Command to get all essential packages 
RUN go get 
# To create the build of our project 
RUN go build -o main .
# Command to run the application
CMD ["/app/main"]
