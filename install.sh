#!/bin/bash

source set-go-path.sh

export DNSPORT=8080
go run src/dns-ep/dns-ep.go
