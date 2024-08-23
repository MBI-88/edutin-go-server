FROM golang:1.22-alpine3.20 as builder 

WORKDIR /app
COPY . .
RUN go mod tidy && \ 
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./web
    

FROM alpine:3.20 as deployment
WORKDIR /app
ARG PORT=80
ENV PORT=${PORT}
COPY --from=builder /app/web ./web
COPY ./data ./data
COPY ./templates ./templates
COPY ./static ./static

RUN addgroup -S appgroup && adduser -S appuser -G appgroup && \
    chown appuser:appgroup ./web && \ 
    chown -R appuser:appgroup ./data ./templates ./static && \
    chmod -R 0775 ./data ./templates ./static

USER appuser   
EXPOSE ${PORT}
 

CMD ["./web"]