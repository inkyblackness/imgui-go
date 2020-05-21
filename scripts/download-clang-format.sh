#!/bin/bash

# This script downloads clang-format as a utility.
# The referenced package is actually an NPM package,
# but why install npm and all that comes with it, when a direct download works as fine.
# Read more about clang-format here: https://clang.llvm.org/docs/ClangFormat.html

PACKAGE_VERSION=1.4.0

wget https://github.com/angular/clang-format/archive/v${PACKAGE_VERSION}.tar.gz
tar -zxvf v${PACKAGE_VERSION}.tar.gz clang-format-${PACKAGE_VERSION}/bin/linux_x64/clang-format --strip-components=3
