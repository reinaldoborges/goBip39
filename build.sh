#!/bin/bash

set -e

DIST_DIR="dist"
EXEC_NAME="bip39"

mkdir -p "$DIST_DIR"
go build -o "$DIST_DIR/$EXEC_NAME" ./cmd

echo "Build complete: $DIST_DIR/$EXEC_NAME"