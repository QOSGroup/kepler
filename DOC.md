# 开发者文档

关键结构体说明

## Subject

* `CN`: 证书的唯一名称，参照OpenSSL的common name

Note: 将来可能增加其他字段

## CertificateSigningRequest

* `IsCa`: 是否可以签发证书
* `Subj`: 主题信息
* `IsBanker`: 是否可以铸币
* `NotBefore`: 最早有效期
* `NotAfter`: 最晚有效期
* `PublicKey`: 请求者的公钥信息

# Issuer

* `PublicKey`: 签发者的公钥
* `Subj`: 签发者的主题信息

## Certificate

* `CSR`: 请求者信息
* `CA`: 签发者信息
* `Signature`: 签名信息
