#!/bin/sh

. ~/.bashrc
rm -f fileserver2 fileserver2.exe
build32
winbuild

ver="$( grep 'version=' *.go | awk '{ print $2 }' FS='=' | \
              awk '{ print $1 }' FS=',' )"
tar -czf fileserver2.v$ver.tar.gz fileserver2 fileserver2.exe

# end.
