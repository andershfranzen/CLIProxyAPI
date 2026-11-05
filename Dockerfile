# Build the management web panel from the named "panel" build context:
#   docker build --build-context panel=/path/to/Cli-Proxy-API-Management-Center ...
FROM oven/bun:1.4.3 AS panel-builder

WORKDIR /panel

COPY --from=panel . .

ARG PANEL_VERSION=

RUN rm -rf node_modules dist && bun install --frozen-lockfile && VERSION="${PANEL_VERSION}" bun run build

FROM golang:1.26-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends build-essential git && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=none
ARG BUILD_DATE=unknown

RUN CGO_ENABLED=1 GOOS=linux go build -buildvcs=false -ldflags="-s -w -X 'main.Version=${VERSION}' -X 'main.Commit=${COMMIT}' -X 'main.BuildDate=${BUILD_DATE}'" -o ./a5proxy ./cmd/server/

FROM debian:bookworm

RUN apt-get update && apt-get install -y --no-install-recommends tzdata ca-certificates && rm -rf /var/lib/apt/lists/*

RUN mkdir /a5proxy

COPY --from=builder ./app/a5proxy /a5proxy/a5proxy

COPY config.example.yaml /a5proxy/config.example.yaml

# The management panel is bundled into the image and never downloaded at runtime.
COPY --from=panel-builder /panel/dist/index.html /a5proxy/static/management.html

ENV MANAGEMENT_STATIC_PATH=/a5proxy/static

WORKDIR /a5proxy

EXPOSE 8317

ENV TZ=Europe/Copenhagen

RUN cp /usr/share/zoneinfo/${TZ} /etc/localtime && echo "${TZ}" > /etc/timezone

CMD ["./a5proxy"]
