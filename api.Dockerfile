FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod verify

COPY . .

RUN go build -o /usr/local/bin/main .

#FROM golang:latest
#
#WORKDIR /app
#
#COPY --from=builder /app/main /app/main

#COPY ./entrypoint.sh /app/entrypoint.sh

 # wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
#ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
#RUN chmod +x /usr/local/bin/wait-for /app/entrypoint.sh

CMD [ "main" ]