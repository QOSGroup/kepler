# 开发者文档

关键结构体说明

## CommonSubject

* `CN`: 证书的唯一名称，参照OpenSSL的common name

## QSCSubject

* `ChainId`: 证书可用链
* `Name`: 币名    
* `Banker`: Banker pubKey

## QCPSubject

* `ChainId`: 证书可用链
* `QCPChain`: 联盟链

## CertificateSigningRequest

* `IsCa`: 是否可以签发证书
* `Subj`: 主题信息
* `NotBefore`: 最早有效期
* `NotAfter`: 最晚有效期
* `PublicKey`: 请求者的公钥信息

# Issuer

* `PublicKey`: 签发者的公钥
* `Subj`: 签发者的主题信息

## Certificate

* `CSR`: 请求者信息(type: CertificateSigningRequest)
* `CA`: 签发者信息(type: Issuer)
* `Signature`: 签名信息(type: []byte)
