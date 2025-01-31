# 创建 dist 目录（如果不存在）
if (!(Test-Path "dist")) {
    New-Item -ItemType Directory -Path "dist" | Out-Null
}

# 编译 Linux 版本
$Env:GOOS = "linux"
$Env:GOARCH = "amd64"
go build -o "dist/webhook_linux" main.go

# 编译 Windows 版本
$Env:GOOS = "windows"
$Env:GOARCH = "amd64"
go build -o "dist/webhook_x64.exe" main.go

Write-Output "Build success"
