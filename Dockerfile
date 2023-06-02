FROM golang:1.17

WORkDIR /usr/src/app

COPY . .

RUN go build -o /server

EXPOSE 3000

CMD ["/server"]