#!/bin/sh

# Script: mock.sh
# Description: Generate mock using mockgen with fixed source/dest directories
# Usage: ./mock.sh <filename.go>
# Example: ./mock.sh permission-repository.go

SOURCE_DIR="internal/core/ports"
DEST_DIR="tests/mock"
PACKAGE_NAME="mock"

# Check if filename is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <filename.go>"
  exit 1
fi

FILENAME="$1"
SOURCE_PATH="${SOURCE_DIR}/${FILENAME}"
DEST_PATH="${DEST_DIR}/mock_${FILENAME}"

# Check if source file exists
if [ ! -f "$SOURCE_PATH" ]; then
  echo "❌ Source file not found: $SOURCE_PATH"
  exit 1
fi

# Ensure destination directory exists
mkdir -p "$DEST_DIR"

echo "🔧 Generating mock from $SOURCE_PATH..."
mockgen -source="$SOURCE_PATH" -destination="$DEST_PATH" -package="$PACKAGE_NAME"

echo "✅ Mock generated at $DEST_PATH"
