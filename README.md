# kepler

[![version](https://img.shields.io/github/tag/QOSGroup/kepler.svg)](https://github.com/QOSGroup/kepler/releases/latest)
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

* `root.pri`: Private Key
* `root.pub`: Public Key
* `root.csr`: Certificate Signing Request
* `root.crt`: Certificate File
* `trust.crts`: Trusted Root Certificate List

## Usage

First you need to create the root certificate, then the union chain certificate, and finally optionally the banker and replay

### ROOT

```
kepler genkey --out-private-key root.pri --out-public-key root.pub
kepler trust --in-public-key root.pub --out-trust-crts trust.crts
kepler req --in-public-key root.pub --is-ca true
kepler sign --in-key-pri root.pri --in-key-pub root.pub

kepler verify
kepler show

```

### QSC

```
kepler genkey --out-private-key qsc.pri --out-public-key qsc.pub $VERBOSE
kepler req --in-public-key qsc.pub --cn QSC --out-sign-req qsc.csr
kepler sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req qsc.csr --out-signed-ca qsc.crt

kepler verify --in-signed-ca qsc.crt
kepler show --in-csr-file qsc.csr --in-crt-file qsc.crt
```

### BANKER

```
kepler genkey --out-private-key banker.pri --out-public-key banker.pub $VERBOSE
kepler req --in-public-key banker.pub --cn QSC --is-banker true --out-sign-req banker.csr
kepler sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req banker.csr --out-signed-ca banker.crt

kepler verify --in-signed-ca banker.crt
kepler show --in-csr-file banker.csr --in-crt-file banker.crt
```

### REPLAY

```
kepler genkey --out-private-key relay.pri --out-public-key relay.pub $VERBOSE
kepler req --in-public-key relay.pub --cn QSC_RELAY --out-sign-req relay.csr
kepler sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req relay.csr --out-signed-ca relay.crt

kepler verify --in-signed-ca relay.crt
kepler show --in-csr-file relay.csr --in-crt-file relay.crt
```

## TODO
 
 - [ ] Two-factor
 - [ ] Add test case
 - [ ] Test illegal certificates
 - [ ] Conversion tools, pem and der to bcm 
 - [ ] Http RPC
 - [ ] Cobra do not use global var

## Acknowledgements

 * [openssl](https://github.com/openssl/openssl)
 * [tendermint](https://github.com/tendermint/tendermint)

## Disclaimer

This is work in progress. Mechanisms and values are susceptible to change.
