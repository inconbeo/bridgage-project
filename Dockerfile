# build stage
FROM golang:alpine AS build-env
RUN apk add git
ADD ./src /src
RUN cd /src && go build -o goapp

# final stage
FROM alpine
RUN adduser -S appuser -G root
WORKDIR /app
RUN chmod -R 775 /app
RUN chown -R appuser:root /app

COPY --from=build-env /src/goapp /app/


USER appuser
ENTRYPOINT ./goapp
