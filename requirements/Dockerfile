FROM golang:1.22.2
WORKDIR /project
COPY project /project
RUN go build -o app main.go
EXPOSE 8080
CMD [ "./app" ]
