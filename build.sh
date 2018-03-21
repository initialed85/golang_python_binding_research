#!/bin/bash

which go 2>/dev/null 1>/dev/null
if [[ $? -ne 0 ]]; then
    echo "error: failed to find go binary- do you have Go 1.9 or Go 1.10 installed?"
    exit 1
fi

GOVERSION=`go version`
if [[ $GOVERSION != *"go1.9"* ]] && [[ $GOVERSION != *"go1.10"* ]]; then
    echo "error: Go version is not 1.9 or 1.10 (was $GOVERSION)"
    exit 1
fi

if [[ "$1" == "" ]]; then
    echo "error: need to specify a go package"
    exit 1
fi

PACKAGE=$1

export GOPATH=`pwd`

export PYTHONPATH=`pwd`/src/github.com/go-python/gopy/

echo "cleaning up output folder"
rm -frv src/$PACKAGE/*.pyc
rm -frv src/$PACKAGE/py2/*.pyc
rm -frv src/$PACKAGE/py2/*.so
rm -frv src/$PACKAGE/cffi/*.pyc
rm -frv src/$PACKAGE/cffi/*.so
rm -frv src/$PACKAGE/cffi/$PACKAGE.py
echo ""

echo "getting assert"
go get github.com/stretchr/testify/assert
echo ""

echo "getting gopy"
go get github.com/go-python/gopy
echo ""

if [[ $GOVERSION == *"go1.10"* ]]; then
    echo "fixing errant pkg-config call in gopy (because we're running Go1.10)"
    sed 's^//#cgo pkg-config: %\[2\]s --cflags --libs^//#cgo pkg-config: %\[2\]s^g' src/github.com/go-python/gopy/bind/gengo.go > temp.go
    mv temp.go src/github.com/go-python/gopy/bind/gengo.go
fi

echo "building gopy"
go build github.com/go-python/gopy
echo ""

echo "building $PACKAGE"
go build -a $PACKAGE
if [[ $? -ne 0 ]]; then
    echo "error: build failed- cannot continue"
    exit 1
fi
echo ""

echo "build $PACKAGE bindings for py2"
./gopy bind -lang="py2" -output="src/$PACKAGE/py2" -symbols=true -work=false $PACKAGE
echo ""

echo "build $PACKAGE bindings for cffi"
./gopy bind -lang="cffi" -output="src/$PACKAGE/cffi" -symbols=true -work=false $PACKAGE
echo ""

TEMP_FILE=$(mktemp)
echo "fixing broken cffi output (this is yuck)"

sed 's/\[\]str):/list):/g' src/$PACKAGE/cffi/$PACKAGE.py > $TEMP_FILE
mv $TEMP_FILE src/$PACKAGE/cffi/$PACKAGE.py

sed 's/\[\]int):/list):/g' src/$PACKAGE/cffi/$PACKAGE.py > $TEMP_FILE
mv $TEMP_FILE src/$PACKAGE/cffi/$PACKAGE.py

sed 's/\[\]long):/list):/g' src/$PACKAGE/cffi/$PACKAGE.py > $TEMP_FILE
mv $TEMP_FILE src/$PACKAGE/cffi/$PACKAGE.py

sed 's/\[\]float):/list):/g' src/$PACKAGE/cffi/$PACKAGE.py > $TEMP_FILE
mv $TEMP_FILE src/$PACKAGE/cffi/$PACKAGE.py

sed 's/\[\]bool):/list):/g' src/$PACKAGE/cffi/$PACKAGE.py > $TEMP_FILE
mv $TEMP_FILE src/$PACKAGE/cffi/$PACKAGE.py
