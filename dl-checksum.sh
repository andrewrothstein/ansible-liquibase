#!/bin/env sh
VER=3.8.0
MIRROR=https://github.com/liquibase/liquibase/releases/download
TGZ=liquibase-${VER}-bin.tar.gz
URL=$MIRROR/liquibase-parent-${VER}/${TGZ}
DIR=~/Downloads
LFILE=$DIR/$TGZ
if [ ! -e $LFILE ]; then
    wget -q -O $LFILE $URL
fi

printf "  # %s\n" $URL
printf "  '%s': sha256:%s\n" $VER `sha256sum $LFILE | awk '{print $1}'`
