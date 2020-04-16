#!/usr/bin/env sh
DIR=~/Downloads

# https://github.com/liquibase/liquibase/releases/download/v3.8.4/liquibase-3.8.4.tar.gz
MIRROR=https://github.com/liquibase/liquibase/releases/download

dl_ver() {
    local ver=$1
    local file=liquibase-${ver}.tar.gz
    local url=$MIRROR/v${ver}/${file}
    local lfile=$DIR/$file
    if [ ! -e $lfile ]; then
        wget -q -O $lfile $url
    fi

    printf "  # %s\n" $url
    printf "  '%s': sha256:%s\n" $ver $(sha256sum $lfile | awk '{print $1}')
}

dl_ver ${1:-3.8.9}
