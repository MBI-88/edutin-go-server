FROM golang:1.22-alpine3.20 as builder 

WORKDIR /app
COPY . .
RUN go mod tidy && \ 
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./server


FROM alpine:3.20 as deployment
WORKDIR /app
COPY ./data .

COPY --from=builder /app/server ./server
RUN addgroup -S appgroup && adduser -S appuser -G appgroup && \
    chown appuser:appgroup ./server

USER appuser   
EXPOSE 80

CMD [ "./server" ]