#!/bin/bash

set -x
set -e

PWD=$(cd "$(dirname "$0")"; pwd)

cd $PWD


CMD=`basename "$PWD"`

$PWD/$CMD genkey
$PWD/$CMD req
$PWD/$CMD sign
$PWD/$CMD verify
