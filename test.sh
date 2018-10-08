#!/bin/bash

# set -x
set -e

PWD=$(cd "$(dirname "$0")"; pwd)
# VERBOSE="-v"

cd $PWD
CMD=`basename "$PWD"`

rm -f $CMD *.pri *.pub *.csr *.crt trust.crts

go build

# ROOT
$PWD/$CMD genkey
$PWD/$CMD trust --in-public-key key.pub --out-trust-crts trust.crts
$PWD/$CMD req
$PWD/$CMD sign
$PWD/$CMD verify
$PWD/$CMD show


# QOS 
$PWD/$CMD genkey --out-private-key qos.pri --out-public-key qos.pub $VERBOSE
$PWD/$CMD req --in-public-key qos.pub --cn QOS --out-sign-req qos.csr
$PWD/$CMD sign  --in-key-pri key.pri --in-key-pub key.pub --in-sign-req qos.csr --out-signed-ca qos.crt
$PWD/$CMD verify --in-signed-ca qos.crt
$PWD/$CMD show --in-csr-file qos.csr --in-crt-file qos.crt


# QSC
$PWD/$CMD genkey --out-private-key qsc.pri --out-public-key qsc.pub $VERBOSE
$PWD/$CMD req --in-public-key qsc.pub --cn QSC --out-sign-req qsc.csr
$PWD/$CMD sign  --in-key-pri key.pri --in-key-pub key.pub --in-sign-req qsc.csr --out-signed-ca qsc.crt
$PWD/$CMD verify --in-signed-ca qsc.crt
$PWD/$CMD show --in-csr-file qsc.csr --in-crt-file qsc.crt
