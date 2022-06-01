FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o /server

EXPOSE 3000 

CMD [ "/server" ]