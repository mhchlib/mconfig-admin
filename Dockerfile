FROM alpine:3.2
WORKDIR /app
ADD mconfig-admin /app
ADD conf/ /app/conf/
ENTRYPOINT [ "./mconfig-admin" ]
