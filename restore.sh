#!/bin/bash

rm -rf tmp/restored
mkdir -p tmp

TIMEFORMAT="elapsed: %Rs"
time ./walg.sh backup-fetch tmp/restored LATEST

cd tmp/restored

find . -type f | sort | xargs md5sum > md5.restored
