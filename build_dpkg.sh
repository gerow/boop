#!/bin/sh

VERSION=1.0-0
BOOP=$GOPATH/src/github.com/gerow/boop

go install github.com/gerow/boop

mkdir -p ./debian/usr/bin
mkdir -p ./debian/etc/init
mkdir -p ./debian/etc/boop
mkdir -p ./debian/DEBIAN

# Apparently this is needed for some versions of debian...
find ./debian -type d | xargs chmod 755

cp $GOPATH/bin/boop ./debian/usr/bin
cp $BOOP/etc/init/boop.conf ./debian/etc/init
cp $BOOP/etc/boop/*.json ./debian/etc/boop
cp $BOOP/etc/control ./debian/DEBIAN

dpkg-deb --build debian

mv debian.deb boop_${VERSION}_amd64.deb

rm -r debian
