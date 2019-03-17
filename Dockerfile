FROM golang
WORKDIR $GOPATH/src/serve
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 3000
CMD ["serve"]
