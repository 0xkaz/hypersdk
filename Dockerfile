FROM golang:1.20-alpine

WORKDIR /myvm
COPY . .

RUN apk update
RUN apk upgrade
RUN apk add vim
RUN apk add bash



#FROM alpine:latest
#FROM golang:1.19-alpine as builder

#RUN apk add git
#WORKDIR /tmp
# RUN export VERSION=1.9.12
# RUN export AVALANCHEGO_PATH=/tmp/avalanchego-v${VERSION}/avalanchego
# RUN export AVALANCHEGO_PLUGIN_DIR=/tmp/avalanchego-v${VERSION}/plugins
# RUN git clone https://github.com/ava-labs/avalanchego.git

# RUN git checkout v${VERSION}
WORKDIR /myvm/examples/weavedbvm

RUN /myvm/examples/weavedbvm/scripts/build.sh 
RUN cp /myvm/examples/weavedbvm/build/* /usr/local/bin/
# RUN go build -o /myvm main.go  && \
#     chmod 777 /myvm

# FROM alpine:latest

# RUN apk update
# RUN apk upgrade

# RUN apk add openssh
# RUN apk add vim
# RUN apk add bash
# RUN apk add runit

# WORKDIR /myvm/
# COPY --from=builder /myvm/build /myvm/build

# EXPOSE 3000
EXPOSE 12352
EXPOSE 12353

# CMD ["/bin/sh", "/myvm/examples/weavedbvm/scripts/run.sh"]
CMD ["/bin/bash", "/myvm/examples/weavedbvm/scripts/run.sh"]
