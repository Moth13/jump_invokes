# Build image
FROM golang:1.19.1-alpine3.16 as base

RUN apk add --no-cache gcc musl-dev &&  mkdir -p /root/go/src/invokes
COPY go.mod /root/go/src/invokes/

WORKDIR /root/go/src/invokes
RUN go mod download

COPY internal/ internal/
COPY cmd/invokes/ cmd/invokes/

ENV GOPATH=/root/go/
ENV GOBIN=/root/go/bin
ENV PATH=$PATH:$GOPATH
ENV PATH=$PATH:$

RUN go install github.com/swaggo/swag/cmd/swag@latest && $GOBIN/swag init -o cmd/invokes/docs -d cmd/invokes,internal -g main.go

WORKDIR /root/go/src/invokes

RUN go build -mod=mod -ldflags="-s -w" -o invokes /root/go/src/invokes/cmd/invokes

# Runtime image
FROM golang:1.19.1-alpine3.16 as runtime

RUN mkdir -p /conf_dir

COPY configs/invokes.yml.template /invokes.yml.template
COPY deployment/docker-entrypoint.sh /usr/bin/docker-entrypoint.sh

COPY --from=base /root/go/src/invokes/invokes /usr/local/bin/

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["/usr/local/bin/invokes", "-conf", "/conf_dir/mh_conf.yml"]

EXPOSE 80

ENV LOG_LEVEL info

ENV DB_ENGINE postgresql
 
ENV DB_PORT 5432
ENV DB_HOST dbserver
ENV DB_PROTO postgres://
ENV DB_USER invokes
ENV DB_PASSWD MYSuperDuperPass
ENV DB_NAME invokes
ENV DB_OPTS ?charset=utf8mb4

ENV CORS_ALLOWED_ORIGINS '*'
ENV CORS_ALLOWED_METHODS 'POST,GET,DELETE'
ENV CORS_ALLOWED_HEADERS 'X-Requested-With,Content-Type'
ENV CORS_MAX_AGE 1728000

ENV PORT 80
ENV CAT_CONF false
