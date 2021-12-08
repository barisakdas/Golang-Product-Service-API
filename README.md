# Golang In Memory Key-Value Store API

[![go-github docker-image (latest SemVer)](https://img.shields.io/github/v/release/google/go-github?sort=semver)](https://hub.docker.com/repository/docker/barisakdas/product-service-api)

The project is an api project that makes product, campaign and order transactions written in golang language.

Standard libraries and third-party packages are used in the project.
```go
go 1.16

require (
github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
github.com/jinzhu/gorm v1.9.16 // indirect
gorm.io/driver/postgres v1.2.3 // indirect
gorm.io/gorm v1.22.4 // indirect
)
```

## Installation

The program can be run easily by downloading the publicly published image on `hub.docker.com`:

```bash
docker pull barisakdas/product-service:latest
```

If you want to run the codes by getting the codes on github, the project can be reached by pulling the repository as follows.

```go
git clone https://github.com/barisakdas/Golang-Product-Service-API.git
```

## Usage

When the API is run, detailed information about all open endpoints can be obtained by sending a request to the `http://localhost:8080/` address first.

```bash
curl http://localhost:8080/
```

The project needs a database to run fully. Therefore, it is necessary to install a postgresql in the deploy environment.Afterwards, the relevant endpoints can be called and the desired operations can be performed from the API.
If you examine the address below, you will see how you can install a postgresql server on docker.
```bash
https://hub.docker.com/_/postgres
```

As soon as the project runs, it will build its own database and tables.
Just having a running postgresql on localhost will be enough.