# Detect OS
UNAME_S := $(shell uname -s 2>/dev/null || echo Windows)

.PHONY: all install build clean help

all: install

install:
ifeq ($(UNAME_S),Windows)
	@echo "Running Windows installation script..."
	@powershell -ExecutionPolicy Bypass -File install.ps1
else
	@echo "Running Linux/macOS installation script..."
	@bash install.sh
endif

build:
ifeq ($(UNAME_S),Windows)
	@echo "[INFO] Building for Windows..."
	@go build -o pjsw.exe main.go
	@echo "[INFO] Build complete: pjsw.exe"
else
	@echo "[INFO] Building for Linux/macOS..."
	@go build -o pjsw main.go
	@echo "[INFO] Build complete: pjsw"
endif

clean:
	@echo "[INFO] Cleaning build artifacts..."
ifeq ($(UNAME_S),Windows)
	@if exist pjsw.exe del pjsw.exe
else
	@rm -f pjsw
endif
	@echo "[INFO] Clean complete"

help:
	@echo "pjsw Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  make install  - Install pjsw (runs platform-specific script)"
	@echo "  make build    - Build pjsw binary only"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make help     - Show this help message"
