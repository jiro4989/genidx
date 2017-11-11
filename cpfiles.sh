#!/bin/bash
# -*- coding: utf-8 -*-

set -eux

tardir=~/Pictures/comic/
cp ./dist/genidx_linux_64-bit/genidx $tardir
cp ./template.html $tardir
