ARG build_image=golang:latest

ARG BUILDPLATFORM
FROM --platform=${BUILDPLATFORM} ${build_image} AS builder

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

COPY . /src
WORKDIR /src
ENV CGO_ENABLED=0
RUN go mod download
RUN go build -o /assets/check ./check
RUN go build -o /assets/in ./in
RUN go build -o /assets/out ./out

FROM builder AS tests
RUN echo "No tests to run for this resource"

FROM scratch AS resource
COPY --from=builder assets/ /opt/resource/
