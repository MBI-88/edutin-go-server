FROM golang:1.19

WORKDIR /app 

COPY go.mod ./
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o edugo

EXPOSE 80

CMD [ "./edugo" ]