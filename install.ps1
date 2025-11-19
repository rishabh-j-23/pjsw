# pjsw Installation Script for Windows
# Run this script in PowerShell

Write-Host "=== pjsw Installation Script for Windows ===" -ForegroundColor Cyan
Write-Host ""

# Check for fzf
try {
    $null = Get-Command fzf -ErrorAction Stop
    Write-Host "[INFO] fzf found" -ForegroundColor Green
} catch {
    Write-Host "Error: fzf is required but not installed." -ForegroundColor Red
    Write-Host "Please install fzf and try again."
    Write-Host "Visit: https://github.com/junegunn/fzf"
    exit 1
}

# Fetch dependencies
Write-Host "[INFO] Fetching dependencies..." -ForegroundColor Yellow
go mod tidy

# Build binary
Write-Host "[INFO] Building pjsw..." -ForegroundColor Yellow
go build -o pjsw.exe main.go

# Create config directory
$ConfigDir = "$env:USERPROFILE\.pjsw"
Write-Host "[INFO] Creating config directory at $ConfigDir" -ForegroundColor Yellow
if (-not (Test-Path $ConfigDir)) {
    New-Item -ItemType Directory -Path $ConfigDir -Force | Out-Null
}

# Install binary
$InstallDir = "$env:USERPROFILE\bin"
Write-Host "[INFO] Installing pjsw to $InstallDir" -ForegroundColor Yellow
if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}
Copy-Item pjsw.exe "$InstallDir\" -Force

Write-Host ""
Write-Host "=== Installation Complete ===" -ForegroundColor Green
Write-Host ""
Write-Host "--------------------------------ACTION REQUIRED----------------------------------" -ForegroundColor Yellow
Write-Host "Add the following directory to your PATH if not already present:" -ForegroundColor White
Write-Host ""
Write-Host "    $InstallDir" -ForegroundColor Cyan
Write-Host ""
Write-Host "To add to PATH permanently, run:" -ForegroundColor White
Write-Host ""
Write-Host '    $env:Path += ";$env:USERPROFILE\bin"' -ForegroundColor Cyan
Write-Host '    [Environment]::SetEnvironmentVariable("Path", $env:Path, "User")' -ForegroundColor Cyan
Write-Host ""
Write-Host "Then restart your terminal to apply the changes" -ForegroundColor White
Write-Host "---------------------------------------------------------------------------------" -ForegroundColor Yellow
