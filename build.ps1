param(
    [switch]$NoPrompt
)

$ErrorActionPreference = "Stop"

$APP_NAME = "gohm"
$BUILD_DIR = "build"

if (-not $NoPrompt) {
    $NewVersion = Read-Host "Enter new version (leave blank to keep current)"
    if ($NewVersion -ne "") {
        $GoModuleContent = $GoModuleContent -replace 'const GOHM_VERSION = "[^"]+"', "const GOHM_VERSION = `"$NewVersion`""
        Set-Content "main.go" -Value $GoModuleContent -NoNewline
        Write-Host "Updated GOHM_VERSION to: $NewVersion" -ForegroundColor Green
    }
}

$GoModuleContent = Get-Content "main.go" -Raw
if ($GoModuleContent -match 'const GOHM_VERSION = "([^"]+)"') {
    $VERSION = $Matches[1]
}
else {
    Write-Error "Could not extract GOHM_VERSION from main.go"
}

Write-Host "Building version: $VERSION" -ForegroundColor Cyan

if (Test-Path $BUILD_DIR) {
    Remove-Item -Recurse -Force $BUILD_DIR
}
New-Item -ItemType Directory -Path $BUILD_DIR | Out-Null

$Platforms = @(
    @{ GOOS = "windows"; GOARCH = "amd64"; Ext = ".exe" },
    @{ GOOS = "windows"; GOARCH = "arm64"; Ext = ".exe" },
    @{ GOOS = "linux"; GOARCH = "amd64"; Ext = "" },
    @{ GOOS = "linux"; GOARCH = "arm64"; Ext = "" },
    @{ GOOS = "darwin"; GOARCH = "amd64"; Ext = "" },
    @{ GOOS = "darwin"; GOARCH = "arm64"; Ext = "" }
)

foreach ($Platform in $Platforms) {
    $Env:GOOS = $Platform.GOOS
    $Env:GOARCH = $Platform.GOARCH
    $OutputName = "$BUILD_DIR/${APP_NAME}-${VERSION}-$($Platform.GOOS)-$($Platform.GOARCH)$($Platform.Ext)"

    Write-Host "  Building for $($Platform.GOOS)/$($Platform.GOARCH)..." -NoNewline

    try {
        go build -ldflags="-s -w" -o $OutputName .
        Write-Host " Done" -ForegroundColor Green
    }
    catch {
        Write-Host " Failed" -ForegroundColor Red
        Write-Host "    Error: $_" -ForegroundColor Red
    }
}

Remove-Item Env:GOOS -ErrorAction SilentlyContinue
Remove-Item Env:GOARCH -ErrorAction SilentlyContinue

Write-Host "`nBuild complete" -ForegroundColor Green
Get-ChildItem $BUILD_DIR | % { Write-Host "  - $($_.Name)" }
