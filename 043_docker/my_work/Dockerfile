FROM golang:alpine
#RUN apt-get update && apt-get -y upgrade
WORKDIR /go/src/app
COPY . .
#ADD . /go/src/github.com/ckpinguin/docker-go-example/mygo
#RUN go install github.com/ckpinguin/docker-go-example/mygo
#ENTRYPOINT [ "/go/bin/mygo" ]
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
CMD ["go-wrapper", "run"] # ["app"]
#EXPOSE 8080
#RUN apt-get -y update && apt-get -y upgrade && apt-get install -y fortune-mod cowsay
#CMD /usr/games/fortune | /usr/games/cowsay