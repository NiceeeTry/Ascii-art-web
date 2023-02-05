    FROM golang:latest
    LABEL maintainer="Alikhan <Alikhanaitbayevpost@gmail.com>"
    RUN mkdir -p /app/
    WORKDIR /app/
    COPY ./ /app/
    RUN go build -o main .
    CMD ["./main"]