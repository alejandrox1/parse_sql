FROM alpine:latest
MAINTAINER Jorge Alarcon
# go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
ADD sql_setup /usr/bin/sql_setup

WORKDIR app/
ADD setup.sql .
ENTRYPOINT ["sql_setup"]
