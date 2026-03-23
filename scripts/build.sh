#!/usr/bin/env bash
# Build script for terraform-provider-cloudflare
# Builds for multiple platforms and creates release artifacts

set -euo pipefail

VERSION="${1:-dev}"
DIST_DIR="${DIST_DIR:-dist}"

mkdir -p "$DIST_DIR"
rm -rf "${DIST_DIR:?}"/*

echo "Building terraform-provider-cloudflare v${VERSION}..."

# Linux AMD64
echo "Building linux/amd64..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags "-s -w -X main.version=${VERSION}" \
    -o "${DIST_DIR}/terraform-provider-cloudflare_${VERSION}_linux_amd64" .

# Linux ARM64
echo "Building linux/arm64..."
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build \
    -ldflags "-s -w -X main.version=${VERSION}" \
    -o "${DIST_DIR}/terraform-provider-cloudflare_${VERSION}_linux_arm64" .

# Darwin AMD64
echo "Building darwin/amd64..."
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags "-s -w -X main.version=${VERSION}" \
    -o "${DIST_DIR}/terraform-provider-cloudflare_${VERSION}_darwin_amd64" .

# Darwin ARM64
echo "Building darwin/arm64..."
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build \
    -ldflags "-s -w -X main.version=${VERSION}" \
    -o "${DIST_DIR}/terraform-provider-cloudflare_${VERSION}_darwin_arm64" .

# Windows AMD64
echo "Building windows/amd64..."
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags "-s -w -X main.version=${VERSION}" \
    -o "${DIST_DIR}/terraform-provider-cloudflare_${VERSION}_windows_amd64.exe" .

echo "All platforms built successfully!"

# Create zip archives
cd "$DIST_DIR"
for binary in terraform-provider-cloudflare_${VERSION}_*; do
    if [[ -f "$binary" ]]; then
        echo "Zipping $binary..."
        zip "${binary}.zip" "$binary"
        rm "$binary"
    fi
done

# Create checksums
echo "Creating checksums..."
sha256sum terraform-provider-cloudflare_${VERSION}_*.zip > terraform-provider-cloudflare_${VERSION}_SHA256SUMS

echo "Build complete! Artifacts in ${DIST_DIR}/"
ls -la
