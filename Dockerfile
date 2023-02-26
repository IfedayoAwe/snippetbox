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
COPY --from=builder /home/snippetbox/snippetbox /home/snippetbox

USER root

# Create the logs directory with the appropriate permissions
RUN mkdir -p /home/snippetbox/logs && chown -R snippetbox:snippetbox /home/snippetbox/logs && chmod 755 /home/snippetbox/logs

USER snippetbox
WORKDIR /home
EXPOSE 4000

ENTRYPOINT ["/home/snippetbox"]