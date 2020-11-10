# Building image containing the application
# NOTES:
# -> Complexity added due to running as non-privileged user

# Base-image for building from source (build)
ARG BUILD_IMAGE

# Base-image for running application (runtime)
ARG RUNTIME_IMAGE

# -----------------------------------------------
# Building binary from source
# -----------------------------------------------
FROM ${BUILD_IMAGE:-golang:alpine} AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV USERNAME=gorun \
    USERID="10001"

WORKDIR $GOPATH/src

# Creating non-privileged user
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home /nonexistent \    
    --shell /sbin/nologin \    
    --no-create-home \    
    --uid ${USERID} \    
    ${USERNAME}

COPY . .

RUN set -e \
    && go build -ldflags="-w -s" -o /go/bin/goapp . \
    && chown -R ${USERNAME}:${USERNAME} /go/bin/ \
    && chmod -R 0500 /go/bin/

# -----------------------------------------------
# Creating minimal image for export
# -----------------------------------------------
FROM ${RUNTIME_IMAGE:-scratch} AS bin

ENV USERNAME gorun

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/goapp /go/bin/goapp

USER gorun:gorun
ENTRYPOINT ["/go/bin/goapp"]
CMD ["--help"]