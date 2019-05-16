#! /bin/bash

set -euo pipefail

GOOS=linux
GOARCH=arm
GOARM=7

# build binary
go build -o bin/arm7-linux/1d-dungeon

# copy to RPI
scp bin/arm7-linux/1d-dungeon $DUNGEON_USER@$DUNGEON_HOST:~/1d-dungeon

# run on RPI
ssh $DUNGEON_USER@$DUNGERON_HOST ./1d-dungeon
