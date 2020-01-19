# go-realworld

![](https://github.com/minchao/go-realworld/workflows/CI/badge.svg?branch=master)

A hexagonal architecture implementation of the realworld example app

## Development

### System requirements

- [Git](https://git-scm.com/)
- [Go 1.13+](https://golang.org/)
- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/) (optional)

### Build

Build realworld app:

```bash
$ make build
```

### Build Docker

Build Docker image:

```bash
$ make docker-image DOCKER_VERSION=latest
```

Run the Docker image:

```bash
$ docker run -it --rm -p 8080:8080 realworld
```

Now you should be able to access API via http://localhost:8080/
