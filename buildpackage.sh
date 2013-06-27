#!/bin/sh

NAME="tomcatinfo"
DESCRIPTION="Shell-based tomcat info gathering tool"
MAINTAINER="James Andariese <james@locationlabs.com>"
VERSION="0.1"

D=`mktemp -d`
if [ $? -ne 0 ];then
    exit
fi
trap "rm -rf \"$D\"" EXIT
mkdir -p "$D/usr/bin"
(cd tomcatinfo;go build)
cp tomcatinfo/tomcatinfo "$D/usr/bin"
fpm -v "$VERSION" -C "$D" -s dir -t deb -n "$NAME" --description "$DESCRIPTION" -m "$MAINTAINER" --deb-user=root --deb-group=root .
