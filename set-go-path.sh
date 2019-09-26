#!/bin/bash

export GOPATH=`pwd`

current_path=`pwd`
cd src/dns-ep/
export GOBIN=`pwd`
cd $current_path
