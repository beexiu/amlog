#!/bin/bash

ROOTDIR=$(cd "$(dirname "$0")"; pwd)
cd $ROOTDIR

$ROOTDIR/build.go api

$ROOTDIR/bin/api
