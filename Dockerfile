# syntax=docker/dockerfile:1

FROM golang:1

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

COPY ./ ./

RUN go install github.com/air-verse/air@latest

CMD [ "air", "-c", ".air.toml" ]
