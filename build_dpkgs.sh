#!/bin/bash

# This script assumes you have cross compilation for go set
# up for amd64 and 386

VERSION=1.0-0
BOOP=$GOPATH/src/github.com/gerow/boop

source $BOOP/etc/bash/crosscompile.bash

mkdir -p ./debian/usr/bin
mkdir -p ./debian/etc/init
mkdir -p ./debian/etc/boop
mkdir -p ./debian/DEBIAN

# Apparently this is needed for some versions of debian...
find ./debian -type d | xargs chmod 755

cp $BOOP/etc/init/boop.conf ./debian/etc/init
cp $BOOP/etc/boop/*.json ./debian/etc/boop
cp $BOOP/etc/control ./debian/DEBIAN


# First build for amd64
go-linux-amd64 build $BOOP/boop/main.go
cp main ./debian/usr/bin/boop
sed s\/%%VERSION%%\/${VERSION}\/g $BOOP/etc/control | sed 's/%%ARCH%%/amd64/g' > ./debian/DEBIAN/control

dpkg-deb --build debian

mv debian.deb boop_${VERSION}_amd64.deb

# Next build for 386
go-linux-386 build $BOOP/boop/main.go
cp main ./debian/usr/bin/boop

sed s\/%%VERSION%%\/${VERSION}\/g $BOOP/etc/control | sed 's/%%ARCH%%/i386/g' > ./debian/DEBIAN/control

dpkg-deb --build debian

mv debian.deb boop_${VERSION}_386.deb

# Next we should build arm, but I haven't
# quite figured out how to do that yet

rm -r debian
rm main
