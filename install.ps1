# Ensure the script is run as Administrator
if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Host "This script must be run as an Administrator. Restarting as administrator..."
    Start-Process powershell.exe "-NoProfile -ExecutionPolicy Bypass -File $($MyInvocation.MyCommand.Path)" -Verb RunAs
    exit
}

# Define directories
$binDir = "bin"
$configDir = "."
$installDir = "C:\Program Files\IITJ Autoproxy"
$configInstallDir = "C:\ProgramData\IITJ Autoproxy"

# Create installation directories if they don't exist
if (-Not (Test-Path $installDir)) {
    New-Item -Path $installDir -ItemType Directory
}

if (-Not (Test-Path $configInstallDir)) {
    New-Item -Path $configInstallDir -ItemType Directory
}

# Copy binaries
Copy-Item -Path "$binDir\*" -Destination $installDir -Force
Copy-Item -Path "$configDir\uninstall.ps1" -Destination $installDir -Force

# Copy example config, README, and LICENSE to config folder
Copy-Item -Path "$configDir\autoproxy.config" -Destination $configInstallDir -Force
Copy-Item -Path "$configDir\README.md" -Destination $configInstallDir -Force
Copy-Item -Path "$configDir\LICENSE" -Destination $configInstallDir -Force

# Add to PATH environment variable
$envPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
if (-Not ($envPath -like "*$installDir*")) {
    [Environment]::SetEnvironmentVariable("Path", $envPath + ";$installDir", [EnvironmentVariableTarget]::Machine)
}

Write-Host "Installation completed successfully!"
