#!/bin/sh

go install github.com/gerow/boop

mkdir -p ./debian/usr/bin
mkdir -p ./debian/etc/init
mkdir -p ./debian/etc/boop
mkdir -p ./debian/DEBIAN

# Apparently this is needed for some versions of debian...
find ./debian -type d | xargs chmod 755

cp $GOPATH/bin/boop ./debian/usr/bin
cp $GOPATH/src/github.com/gerow/boop/etc/init/boop.conf ./debian/etc/init
cp $GOPATH/src/github.com/gerow/boop/etc/boop/*.json ./debian/etc/boop
cp $GOPATH/src/github.com/gerow/boop/etc/control ./debian/DEBIAN

dpkg-deb --build debian

mv debian.deb boop_1.0-0_amd64.deb

rm -r debian
