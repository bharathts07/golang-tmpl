FROM golang:1.12.5

ARG GITHUB_TOKEN
ARG SERVICE_NAME
ARG VERSION

WORKDIR /tmp/${SERVICE_NAME}

COPY . ./

COPY go.* ./

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v \
	-ldflags="-w -s -X main.version=${VERSION} -X main.serviceName=${SERVICE_NAME}" \
	./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=0 /go/bin/server /bin/server

CMD ["/bin/server"]
