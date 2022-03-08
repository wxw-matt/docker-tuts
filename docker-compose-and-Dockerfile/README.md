# Docker compose and Dockerfile experiment

## Overiride compose file used
`docker-compose.override.yml` is introduced to reused config from `docker-compose.yml`. Docker-compose merges the config parameters from `docker-compose.override.yml` and `docker-compose.yml` to produce the target config file. The parameters in `docker-compose.override.yml` will override the ones in `docker-compose.yml`.

## Build images using arbitrary names for Dockerfile
```yaml
services:
  dev:
    profiles: ['dev']
    build:
      context: ./
      dockerfile: Dockerfile-dev
```
The `Dockerfile` can be named arbitrarily.

## Profiles is used (see the previous example) to allow users to select the services to run


## Multi-stage builds used to reduce the size of production images
See the differences between `Dockerfile-dev` and `Dockerfile-pro`.
Both use `multi-stage builds`, but `Dockerfile-pro` uses `FROM scratch` instead of `FROM alpine:3.15`. The former is smaller than latter.

Besids, as at runtime, `go` programs do not reply on `go` language components, the second stage does not contain `go`-related files. It results in much smaller images.

## Start containers
See `run-dev.sh` and `run-pro.sh`.

## References
1. [docker-compose.yml extends](https://docs.docker.com/compose/extends/)
2. [Profiles](https://docs.docker.com/compose/profiles/)
3. [Multi-stage builds](https://docs.docker.com/develop/develop-images/multistage-build/)
