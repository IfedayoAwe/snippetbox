FROM golang:1.19-alpine AS builder

RUN /sbin/apk update && \
	/sbin/apk --no-cache add ca-certificates git tzdata && \
	/usr/sbin/update-ca-certificates

RUN adduser -D -g '' snippetbox
WORKDIR /home/snippetbox

COPY . /home/snippetbox

ARG VERSION

RUN CGO_ENABLED=0 go build -a -tags netgo,osusergo \
    -ldflags "-extldflags '-static' -s -w" \
    -ldflags "-X main.version=$VERSION" -o snippetbox ./cmd/web

FROM busybox:musl

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /home/snippetbox/snippetbox /home/snippetbox/snippetbox

RUN chown -R snippetbox:snippetbox /home/snippetbox

USER snippetbox
WORKDIR /home/snippetbox
EXPOSE 4000

ENTRYPOINT ["/home/snippetbox/snippetbox"]