## EN | [RU](https://github.com/ghostemeow/goshikimori/blob/master/README.md)

### About
A small library for interacting with shikimori, written in golang.
* Work with API occurs through `OAuth2`. \
And you have to start by familiarizing yourself
with the documentation [first steps](https://github.com/ghostemeow/goshikimori/blob/master/examples/first_steps/README_en.md).
* No dependencies on other libraries.
* The [GNU make](https://www.gnu.org/software/make/manual/make.html)
utility is used for tests and builds.

### Install
```bash
go get github.com/ghostemeow/goshikimori
```

### Examples
* [Click her](https://github.com/ghostemeow/goshikimori/tree/master/examples)

### GODOC / GOPKG documentation
**Godoc** support is also available.

Or you can use the page from the official Go pkg
[website](https://pkg.go.dev/github.com/ghostemeow/goshikimori).\
**P.S.** documentation is late in updating.
```bash
# Method #1: Use docker.
git clone git@github.com:ghostemeow/goshikimori.git && cd goshikimori
make docker

# Open in browser.
http://localhost:1337/pkg/github.com/ghostemeow/goshikimori/
```
```bash
# Method #2(Linux): Install godoc.
go install -v golang.org/x/tools/cmd/godoc@latest
# Add 'export' to the file /home/$USER/.profile and reboot.
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# Check that the application is working properly.
godoc --help

# After installation or if 'godoc' is already installed.
git clone git@github.com:ghostemeow/goshikimori.git && cd goshikimori
make doc

# Open in browser.
http://localhost:1337/pkg/github.com/ghostemeow/goshikimori
```

### Shikimori documentation
* [GraphQL](https://shikimori.io/api/doc/graphql)
* [API v1](https://shikimori.io/api/doc/1.0)
* [API v2](https://shikimori.io/api/doc/2.0)

### Feedback
* Private message on the [website](https://shikimori.io/arctica).
* Open [issue](https://github.com/ghostemeow/goshikimori/issues).
