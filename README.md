# kepler
**Note: Requires Go 1.11+**

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
