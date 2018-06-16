FROM alpine

RUN apk add --no-cache libc6-compat

ENV VERSION="v0.1"

WORKDIR /app
RUN wget -q https://github.com/hunterlong/balancebadge/releases/download/$VERSION/balancebadge-linux-x64
RUN chmod +x balancebadge-linux-x64 && mv balancebadge-linux-x64 /usr/local/bin/balancebadge

EXPOSE 9090

VOLUME /app

ENTRYPOINT balancebadge