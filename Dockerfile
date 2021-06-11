FROM busybox:latest
 #golang:1.16

# Set the Current Working Directory inside the container
#WORKDIR $GOPATH/src/github.com/Smart-Biz-Cloud-Solutions/service-template-golang

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY service-template-golang .

RUN chmod 744 service-template-golang
# Download all the dependencies
#RUN go get -d -v ./...

# Install the package
#RUN go install -v ./...

# This container exposes port 8080 to the outside world
# EXPOSE 8080

# Run the executable
CMD ["./service-template-golang"]