#!/bin/bash

# set -x
set -e

PWD=$(cd "$(dirname "$0")"; pwd)
# VERBOSE="-v"

cd $PWD
CMD=`basename "$PWD"`

rm -f $CMD *.pri *.pub *.csr *.crt trust.crts

go build


########
# ROOT #
########

$PWD/$CMD genkey --out-private-key root.pri --out-public-key root.pub
$PWD/$CMD trust --in-public-key root.pub --out-trust-crts trust.crts
$PWD/$CMD req --in-public-key root.pub --is-ca true
$PWD/$CMD sign --in-key-pri root.pri --in-key-pub root.pub
sleep 1
$PWD/$CMD verify
$PWD/$CMD show

#######
# QSC #
#######

$PWD/$CMD genkey --out-private-key qsc.pri --out-public-key qsc.pub $VERBOSE
$PWD/$CMD req --in-public-key qsc.pub --cn QSC --out-sign-req qsc.csr
$PWD/$CMD sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req qsc.csr --out-signed-ca qsc.crt
sleep 1
$PWD/$CMD verify --in-signed-ca qsc.crt
$PWD/$CMD show --in-csr-file qsc.csr --in-crt-file qsc.crt

##########
# BANKER #
##########

$PWD/$CMD genkey --out-private-key banker.pri --out-public-key banker.pub $VERBOSE
$PWD/$CMD req --in-public-key banker.pub --cn QSC --is-banker true --out-sign-req banker.csr
$PWD/$CMD sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req banker.csr --out-signed-ca banker.crt
sleep 1
$PWD/$CMD verify --in-signed-ca banker.crt
$PWD/$CMD show --in-csr-file banker.csr --in-crt-file banker.crt

##########
# REPLAY #
##########

$PWD/$CMD genkey --out-private-key relay.pri --out-public-key relay.pub $VERBOSE
$PWD/$CMD req --in-public-key relay.pub --cn QSC_RELAY --out-sign-req relay.csr
$PWD/$CMD sign  --in-key-pri root.pri --in-key-pub root.pub --in-sign-req relay.csr --out-signed-ca relay.crt
sleep 1
$PWD/$CMD verify --in-signed-ca relay.crt
$PWD/$CMD show --in-csr-file relay.csr --in-crt-file relay.crt


