# kepler

[![Build Status](https://travis-ci.org/QOSGroup/kepler.svg?branch=master)](https://travis-ci.org/QOSGroup/kepler)
[![codecov](https://codecov.io/gh/QOSGroup/kepler/branch/master/graph/badge.svg)](https://codecov.io/gh/QOSGroup/kepler)
[![Go version](https://img.shields.io/badge/go-1.11.0-blue.svg)](https://github.com/moovweb/gvm)
[![license](https://img.shields.io/github/license/QOSGroup/kepler.svg)](https://github.com/QOSGroup/kepler/blob/master/LICENSE)
[![](https://tokei.rs/b1/github/QOSGroup/kepler?category=lines)](https://github.com/QOSGroup/kepler)

Another certificate format `BCM`, similar to `PEM` OR `DER`

## OpenSSL

* PEM–Format
  
  DER is a binary format for data structures described by ASN.1.

* DER–Format

  Privacy Enhanced Mail (PEM) is a format with goal to embed binary content into a content typed 7bits ASCII format.
  
## The file suffix

* `key.pri`: Private Key
* `key.pub`: Public Key
* `root.csr`: Certificate Signing Request
* `root.crt`: Short For Certificate
* `trust.crts`: Trusted Root Certificate List


## TODO
 
 - [ ] Four-factor
 - [ ] Add test case
 - [ ] Test illegal certificates
 - [ ] Conversion tools, pem and der to bcm 
 - [ ] Http RPC
 - [ ] Cobra do not use global var

## Acknowledgements

 * [tendermint](https://github.com/tendermint/tendermint)
 * [openssl](https://github.com/openssl/openssl)

## Disclaimer

This is work in progress. Mechanisms and values are susceptible to change.
