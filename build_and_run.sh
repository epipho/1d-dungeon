#! /bin/bash

set -euo pipefail

CGO_ENABLED=0 go build -o bin/1d-dungeon
./bin/1d-dungeon $*
