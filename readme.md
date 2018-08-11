# Serve
[![Go Report Card](https://goreportcard.com/badge/github.com/fronbasal/serve)](https://goreportcard.com/report/github.com/fronbasal/serve)
[![Maintainability](https://api.codeclimate.com/v1/badges/fccfa711a959c2efef3d/maintainability)](https://codeclimate.com/github/fronbasal/serve/maintainability)
[![codebeat badge](https://codebeat.co/badges/6aabbecf-5cdf-4ef3-b751-db975182dbcb)](https://codebeat.co/projects/github-com-fronbasal-serve-master)


Serve is a simple CLI to serve data via http.

Use of serve is intended for development purposes only!

## Installation
```bash
go get -v -u github.com/fronbasal/serve
``` 

Binary releases are available on the [release page](https://github.com/fronbasal/serve/releases) or can be compiled utilizing the Makefile.

## Usage
```
usage: serve [<flags>] [<directory>]

Flags:
      --help       Show context-sensitive help (also try --help-long and --help-man).
  -p, --port=3000  The port of the HTTP server.
  -v, --verbose    Enable verbose output

Args:
  [<directory>]  The directory to serve
```

## Maintainers

- Daniel Malik (mail@fronbasal.de)

## License

MIT
