FROM node:14-alpine3.14 as node-build-stage
WORKDIR /app/front
COPY ./front/package*.json ./

# RUN npm install

RUN apk --no-cache --virtual build-dependencies add \
        python3 \
        py3-pip \
        make \
        g++ \
&& npm install \
&& apk del build-dependencies


COPY ../front .
RUN npm run build

FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

#RUN apk update && apk add git
RUN apk add build-base

COPY ../go.mod go.sum ./
RUN go mod download

# only for development
# RUN go get -u github.com/githubnemo/CompileDaemon
# RUN go install github.com/githubnemo/CompileDaemon

ADD .. /app/

# Install libwebp
RUN apk update && \
    apk upgrade -U && \
    apk add libwebp-dev && \
    rm -rf /var/cache/*

RUN GOOS=linux go build ./cmd/hatcomment

### PROD
FROM alpine:3.18 as go-prod
WORKDIR /app
COPY --from=builder /app/hatcomment /app/hatcomment
COPY --from=builder /app/web /app/web
COPY --from=node-build-stage /app/dist /app/dist
# Install libwebp
RUN apk update && \
    apk upgrade -U && \
    apk add libwebp-dev && \
    rm -rf /var/cache/*

RUN mkdir -p ./storage/upload

## Healthcheck for PROD
#RUN apk --no-cache add curl
#HEALTHCHECK --interval=5s --timeout=3s --start-period=10s --retries=3 \
#    CMD curl -fs http://localhost:$PORT/ || exit 1

ENTRYPOINT ./hatcomment

