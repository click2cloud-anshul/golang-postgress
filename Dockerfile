FROM golang:alpine


# Move to working directory /build
WORKDIR /gopost
RUN apk add git
RUN go get -u github.com/gorilla/mux
ADD . /gopost
# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/gopost/main"]