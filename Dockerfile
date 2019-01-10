FROM alpine:latest
RUN apk add --no-cache imagemagick bash
RUN mkdir -p /app && mkdir -p /app/static && mkdir -p /app/templates && mkdir -p /app/tmp
ADD static /app/static
ADD templates /app/templates
ADD web /app/web
WORKDIR /app
EXPOSE 5000
ENTRYPOINT ["/app/web"]
