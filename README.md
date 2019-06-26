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

First you need to create the root certificate, then the union chain certificate, and finally optionally the banker and relay

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
kepler genkey --out-private-key qsc.pri --out-public-key qsc.pub
kepler genkey --out-private-key banker.pri --out-public-key banker.pub
kepler req-qsc --in-public-key qsc.pub --chain-id qos-test --name qstars-test --banker banker.pub --out-sign-req qsc.csr
kepler sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req qsc.csr --out-signed-ca qsc.crt

kepler verify --in-signed-ca qsc.crt
kepler show --in-csr-file qsc.csr --in-crt-file qsc.crt
```

### QCP

```
kepler genkey --out-private-key qcp.pri --out-public-key qcp.pub
kepler req-qcp --in-public-key qcp.pub --chain-id qos-test --qcp-chain qstars-test --out-sign-req qcp.csr
kepler sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req qcp.csr --out-signed-ca qcp.crt

kepler verify --in-signed-ca qcp.crt
kepler show --in-csr-file qcp.csr --in-crt-file qcp.crt
```

### RELAY

```
kepler genkey --out-private-key relay.pri --out-public-key relay.pub
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
