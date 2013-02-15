#!/bin/sh

go install github.com/gerow/boop
cp -r $GOPATH/src/github.com/gerow/boop/etc/boop /etc
cp $GOPATH/src/github.com/gerow/boop/etc/init/boop.conf /etc/init
