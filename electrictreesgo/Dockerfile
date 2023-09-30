# Needs to be Debian rather than Alpine here because of bugs in Go plugin (shared library)
# building in ARM64 alpine, e.g. https://github.com/docker-library/golang/issues/402
# Also you need to use binutils-gold on alpine to even get the code to link but then
# running the linter causes a segmentation fault.  Debian, on the other hand appears to
# work fine :)
FROM golang:1.20.4-bookworm as build
ARG APP_NAME
ARG TARGETARCH
ARG TARGETOS

# Need to convert the TARGETARCH to an environment variable so that the
# Makefile can pick it up and use to build the executable
ENV APP_NAME=${APP_NAME}
ENV TARGETARCH=${TARGETARCH}
ENV TARGETOS=${TARGETOS}

WORKDIR /usr/src/app
COPY ./cmd ./cmd/
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/
COPY go.sum go.mod ./
COPY Makefile VERSION ./

# In the repo this already exists but copying it over would be silly
RUN mkdir ./build

# This downloads dependencies and builds Go tools
RUN make install_tools

# Check that code is working in this context
RUN make generate lint test

# Now build the code
RUN make build


FROM alpine:3.18.4 as deploy
ARG APP_NAME
ARG TARGETARCH
ARG TARGETOS

# These need to be environment variables as well as args so
# they are available at runtime
ENV APP_NAME=${APP_NAME}
ENV TARGETARCH=${TARGETARCH}
ENV TARGETOS=${TARGETOS}

COPY --from=build /usr/src/app/build/${APP_NAME}-${TARGETOS}-${TARGETARCH} /app/${APP_NAME}

# Need to use the shell to run the app so that it can parse the
# environment variables from above.  Also the `exec` allows that
# shell to be completely replaced by the application (saves resources)
CMD [ "/bin/sh", "-c", "exec /app/${APP_NAME}" ]
