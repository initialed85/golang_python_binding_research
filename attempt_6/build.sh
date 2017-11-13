#!/usr/bin/env bash

# export GOROOT=/usr/local/Cellar/go/1.5.4
# echo "GOROOT is $GOROOT"
# echo ""

export GOPATH=`pwd`
echo "GOPATH is $GOPATH"
echo ""

export GO=`which go`
echo "GO is $GO"
echo ""

GOVERSION=`$GO version`
if [[ $GOVERSION != *"go1.9"* ]]; then
    echo "error: Go version is not 1.9 (was $GOVERSION)"
    exit 1
fi

# echo "getting gopy"
# $GO get -v github.com/go-python/gopy
# echo ""

PYTHONPATH="`pwd`/src/github.com/go-python/gopy/"

# echo "building gopy"
# $GO build -a github.com/go-python/gopy
# echo ""

echo "build gosnmp_python bindings"
./gopy bind -lang="py2" -output="attempt_6" github.com/initialed85/attempt_6
echo ""
