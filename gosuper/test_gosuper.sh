#!/bin/bash
set -e
go test -v ./gosuper
go test -v ./supermodels
go test -v ./config
