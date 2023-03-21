#Dockerfile
#From which image we want to build. This is basically our environment.
FROM golang:1.17.6-alpine as Build

#This will copy all the files in our repo to the inside the container at root location.
COPY . . 

#build our binary at root location. Binary name will be main. We are using go modules so gpath env variable should be empty.
RUN GOPATH= go build -o /main main.go 

#we tell docker what to run when this image is run and run it as executable.
ENTRYPOINT [ "/main" ]
