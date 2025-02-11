#adding a builder for multistage build
FROM golang:1.22.2-alpine As builder 

#Metadata 
LABEL owner="soubaaiss" \
      owner="mlarbi" \
      owner="ebelfkih" \
      project="ascii-art-web" \
      version="1.0" \
      description="This is a simple web application that converts text to ASCII art." \
      environment="zone 01"

EXPOSE 8080

# WORKDIR /project
WORKDIR /build

COPY  project .

RUN go build -o main . 


FROM alpine:latest 
WORKDIR /app

#install the bash
RUN apk add --no-cache bash

#copy the builder into the seconde stage

COPY --from=builder /build .

CMD [ "/app/main" ]
