FROM mcr.microsoft.com/devcontainers/go:0-1-bullseye
LABEL one.brotherday.vendor "Brother Day"
LABEL one.brotherday.image.authors "Steve Huguenin <steve@brotherday.one>"
LABEL org.opencontainers.image.source https://github.com/brotherday/mainstream-backend
LABEL version 1.0
LABEL description "Multi-stage Docker image for mainstream backend"
LABEL license AGPL-3.0
WORKDIR /mainstream
ADD . .
RUN go build
FROM scratch
EXPOSE 3000
COPY --from=0 /mainstream/ .
CMD /backend
