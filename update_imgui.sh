#!/usr/bin/env bash

set -e

# This is a small script to fetch the imgui source

TMP_IMGUI=.tmp_imgui
DEST=imgui
UPSTREAM_GIT_URI=https://github.com/ocornut/imgui.git

if [ $# -ne 1 ]; then
  echo "Usage: $0 <git ref>"
  exit 1
fi

GIT_REV=$1

mkdir -p $DEST
echo "Cloning imgui to $TMP_IMGUI"
git clone --depth=1 --branch $GIT_REV $UPSTREAM_GIT_URI $TMP_IMGUI

mkdir -p $DEST
rm -rf $DEST/*

echo "Copying files"
# Copy core files
cp $TMP_IMGUI/*.{h,cpp} $DEST/
echo "package cgo${DEST}" > $DEST/govendorkeep.go

# Copy freetype
mkdir -p $DEST/misc/freetype
cp $TMP_IMGUI/misc/freetype/*.{h,cpp} $DEST/misc/freetype
echo "package cgofreetype" > $DEST/misc/freetype/govendorkeep.go

# Copy license
cp $TMP_IMGUI/LICENSE.txt _licenses/imgui-LICENSE.txt

# Clean up
echo "Removing $TMP_IMGUI"
rm -fr $TMP_IMGUI
