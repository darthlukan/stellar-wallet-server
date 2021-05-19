# Stellar Wallet Server
[![CodeQL](https://github.com/darthlukan/stellar-wallet-server/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/darthlukan/stellar-wallet-server/actions/workflows/codeql-analysis.yml)

Author: Brian Tomlinson <darthlukan@gmail.com>


## Description

Stellar Wallet Server is yet another cryptocurrency Wallet built specifically
with the [Stellar Network](https://www.stellar.org/) in mind, for now. It is meant
to provide a REST API for any number of clients seeking a simple means of interacting
with [Stellar Lumens (XLM)](https://www.stellar.org/lumens).

*IMPORTANT*: _This is a personal project that is not (yet) meant for production use and should be considered highly volatile._


## Deploy

### OpenShift

Stellar Wallet Server currently targets OpenShift 4.X for deployment, the requisite manifests reside in the
`deploy/openshift/` directory of this repository.

```
$ oc apply -f deploy/openshift/
```


## Develop

Development of Stellar Wallet Server requires [Go](https://golang.org/) version 1.16+. New functionality should be provided via packages
containing their own tests. Packages should export a public method to the `api` package and link that handler function
to the [Gin](https://gin-gonic.com/) router in `stellar-wallet-server.go`.


## Test

Tests for Stellar Wallet Server can be executed via the `Makefile`:
```
$ make test
```

The `go test` command will run and provide relevant output.


## Contribute

Contributions from the community are very welcome. Please fork this repository and submit your changes to the `main`
branch via Pull Request. Pull Requests should correspond to an Issue, if one does not exist, please create one before
submitting your PR.

Only Pull Requests which pass the [CodeQL](https://codeql.github.com/docs/codeql-overview/) at a minimum will be
accepted. 


## License

MIT, see LICENSE file.
