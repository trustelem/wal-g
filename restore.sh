#!/bin/bash

rm -rf tmp/restored
mkdir -p tmp

./walg.sh backup-fetch tmp/restored LATEST
