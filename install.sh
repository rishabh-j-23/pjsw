#!/bin/bash

set -e

echo "=== pjsw Installation Script for Linux/macOS ==="
echo ""

# Check for fzf
if ! command -v fzf &> /dev/null; then
    echo "Error: fzf is required but not installed."
    echo "Please install fzf and try again."
    echo "Visit: https://github.com/junegunn/fzf"
    exit 1
fi

echo "[INFO] fzf found"

# Fetch dependencies
echo "[INFO] Fetching dependencies..."
go mod tidy

# Build binary
echo "[INFO] Building pjsw..."
go build -o pjsw main.go

# Create config directory
CONFIG_DIR="$HOME/.pjsw"
echo "[INFO] Creating config directory at $CONFIG_DIR"
mkdir -p "$CONFIG_DIR"

# Copy shell script
echo "[INFO] Copying pjsw-switch.sh to $CONFIG_DIR"
cp ./scripts/pjsw-switch.sh "$CONFIG_DIR/"

# Install binary
INSTALL_DIR="/usr/local/bin"
echo "[INFO] Installing pjsw to $INSTALL_DIR (requires sudo)"
sudo cp pjsw "$INSTALL_DIR/"
sudo chmod +x "$INSTALL_DIR/pjsw"

echo ""
echo "=== Installation Complete ==="
echo ""
echo "--------------------------------ACTION REQUIRED----------------------------------"
echo "Add the following line to your shell rc file (~/.bashrc, ~/.zshrc, etc.):"
echo ""
echo "    source ~/.pjsw/pjsw-switch.sh"
echo ""
echo "Then restart your terminal or run: source ~/.bashrc (or ~/.zshrc)"
echo "---------------------------------------------------------------------------------"
