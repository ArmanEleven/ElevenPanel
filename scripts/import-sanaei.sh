#!/usr/bin/env bash
set -euo pipefail

UPSTREAM_REPO="https://github.com/MHSanaei/3x-ui.git"
UPSTREAM_TAG="v3.9.0"
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

git clone --depth 1 --branch "$UPSTREAM_TAG" "$UPSTREAM_REPO" "$TMP_DIR/3x-ui"

rm -rf .upstream/3x-ui
mkdir -p .upstream/3x-ui
cp -a "$TMP_DIR/3x-ui/." .upstream/3x-ui/

echo "Imported Sanaei/3x-ui $UPSTREAM_TAG into .upstream/3x-ui"
echo "Review license/attribution before promoting files into the product tree."
