all: build install

check-fzf:
	@command -v fzf >/dev/null 2>&1 || { echo >&2 "fzf is required but not installed. Please install it and try again."; exit 1; }

build:
	@echo "[INFO] Checking fzf"
	@make check-fzf
	@echo "[INFO] Fetching dependencies"
	@go mod tidy
	@echo "[INFO] Building pjsw"
	@go build -o pjsw main.go
	@echo "[INFO] Copying pjsw-switch.sh to ~/.pjsw"
	@cp ./scripts/pjsw-switch.sh $(HOME)/.pjsw
	@echo "[INFO] pjsw built"

run:
	go run main.go

install:
	@echo "[INFO] Checking fzf"
	@make check-fzf
	@echo "[INFO] Fetching dependencies"
	@go mod tidy
	@echo "[INFO] installing pjsw"
	@go install pjsw
	@echo "[INFO] pjsw installed"
	@echo "--------------------------------ACTION REQUIRED----------------------------------"
	@echo "[WARN] add 'source ~/.pjsw/pjsw-switch.sh' to the rc file"
	@echo "[INFO] restart your terminal to apply the changes"
	@echo "---------------------------------------------------------------------------------"
