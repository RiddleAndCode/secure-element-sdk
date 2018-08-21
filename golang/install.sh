#!/bin/sh

# shell script for installing the go cryptoauthlib

# move the cryptoauth c library in place
sudo cp ../bin/libcryptoauth.so /usr/lib

# move the dummy go sources in place
cp -R src $GOPATH/.

# move the go binary in place
cp -R pkg $GOPATH/.

