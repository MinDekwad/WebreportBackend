FROM golang:1.16-alpine AS build-src

WORKDIR /src
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tmds-report ./main.go


FROM alpine:3.13.2
# ARG ENV
RUN apk --no-cache add ca-certificates tzdata curl
ENV TZ=Asia/Bangkok
WORKDIR /app
COPY --from=build-src /src/tmds-report .
#
COPY --from=build-src /src/cert/prod.crt /usr/local/share/ca-certificates/prod.crt
COPY --from=build-src /src/cert/prod.crt /app/cert/prod.crt
#
RUN update-ca-certificates

# RUN mkdir -p /app/conf
# COPY --from=build-src /src/conf/example.yaml /app/conf/default.yaml

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN chown -R appuser:appgroup ./* && ls -lah ./cert

USER appuser

CMD [ "./tmds-report" ]
