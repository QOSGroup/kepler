#!/bin/bash

set -x
set -e

PWD=$(cd "$(dirname "$0")"; pwd)

cd $PWD
CMD=`basename "$PWD"`

rm -f $CMD *.pri *.pub *.csr *.crt

go build

$PWD/$CMD genkey
$PWD/$CMD req
$PWD/$CMD sign
$PWD/$CMD verify
