# kepler

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


## TODO
 
 - [ ] Conversion tools, pem and der to bcm 
 - [ ] Http RPC
 - [ ] Cobra do not use global var
 - [ ] Add test case
 - [ ] Test illegal certificates
 - [ ] CA trust list
