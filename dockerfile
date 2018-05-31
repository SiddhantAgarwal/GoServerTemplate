FROM golang

EXPOSE 8080
ENV webserver_path /go/src/github.com/SiddhantAgarwal/GoServerTemplate/
ENV PATH $PATH:$webserver_path
ADD . /go/src/github.com/SiddhantAgarwal/GoServerTemplate

WORKDIR $webserver_path
RUN go build -o myapp . 
ENTRYPOINT ["./myapp"] 