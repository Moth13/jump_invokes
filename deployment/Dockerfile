# Build image
FROM golang:1.16.5-alpine3.13 as base

RUN apk add --no-cache gcc musl-dev

RUN mkdir -p /root/go/src/app

COPY go.mod /root/go/src/app/
RUN cd /root/go/src/app/

WORKDIR /root/go/src/app
RUN go mod download

COPY internal/ internal/
COPY cmd/app/ cmd/app/

ENV GOPATH=/root/go/
ENV GOBIN=/root/go/bin
ENV PATH=$PATH:$GOPATH
ENV PATH=$PATH:$

RUN mkdir -p /root/go/src/app/view/html
RUN mkdir -p /root/go/src/app/view/css
COPY view/html/404.html /root/go/src/app/view/html/404.html 
COPY view/css/404.css /root/go/src/app/view/css/404.css 

WORKDIR /root/go/src/app/cmd/app/

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN $GOBIN/swag init --parseDependency -o /root/go/src/app/docs

WORKDIR /root/go/src/app

RUN go build -mod=mod -ldflags="-s -w" -o app /root/go/src/app/cmd/app

# Runtime image
FROM golang:1.16.5-alpine3.13 as runtime

RUN mkdir -p /conf_dir

COPY configs/app.yml.template /app.yml.template
COPY deployment/docker-entrypoint.sh /usr/bin/docker-entrypoint.sh

RUN mkdir -p /root/go/src/app/view/html
RUN mkdir -p /root/go/src/app/view/css
RUN mkdir /css
COPY view/html/404.html /root/go/src/app/view/html/404.html
COPY view/css/404.css /root/go/src/app/view/css/404.css
COPY view/css/404.css /css

COPY --from=base /root/go/src/app/app /usr/local/bin/

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["/usr/local/bin/app", "-conf", "/conf_dir/mh_conf.yml"]

EXPOSE 80

ENV CORS_ALLOWED_ORIGINS '*'
ENV CORS_ALLOWED_METHODS 'POST,GET,DELETE'
ENV CORS_ALLOWED_HEADERS 'X-Requested-With,Content-Type'
ENV CORS_MAX_AGE 1728000

ENV DEBUG false
ENV PORT 80
ENV PROJET_REVISION_BASE_URL http://localhost:8080
ENV PROJET_BASE_PATH ""
ENV CAT_CONF false
ENV VERSION_PREFIXED_ROUTES true
