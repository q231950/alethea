# alethea

[![CircleCI](https://circleci.com/gh/q231950/alethea.svg?style=svg)](https://circleci.com/gh/q231950/alethea)

## Install

Get the package and install it with `go install`:

```bash
$ go get github.com/q231950/alethea
$ cd $GOPATH/src/github.com/q231950/
$ go install
```

### Postgres

**alethea** uses the following environment variables to store incidents in a Postgres database. You need a running Postgres database and export the following variables:
- `ALETHEA_POSTGRESQL_USER`
- `ALETHEA_POSTGRESQL_PASSWORD`
- `ALETHEA_POSTGRESQL_DATABASE`
- `ALETHEA_POSTGRESQL_ADDRESS`
- `ALETHEA_POSTGRESQL_PORT`

## Run

You can run **alethea** with a single parameter, the port which it will serve. If no port is specified, it will serve `8080`.

```bash
alethea --port=8080
```

Note: Heroku requires to use the exported `PORT` environment variable to be used for services. If a `PORT` environment variable is exported, it will be used to serve **alethea**.