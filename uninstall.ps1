# Ensure the script is run as Administrator
if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Host "This script must be run as an Administrator. Restarting as administrator..."
    Start-Process powershell.exe "-NoProfile -ExecutionPolicy Bypass -File $($MyInvocation.MyCommand.Path)" -Verb RunAs
    exit
}

# Define directories
$installDir = "C:\Program Files\IITJ Autoproxy"
$configInstallDir = "C:\ProgramData\IITJ Autoproxy"

# Remove binaries
if (Test-Path $installDir) {
    Remove-Item -Path "$installDir\*" -Force
    Remove-Item -Path $installDir -Force
}

# Remove config files, README, and LICENSE
if (Test-Path $configInstallDir) {
    Remove-Item -Path "$configInstallDir\*" -Force
    Remove-Item -Path $configInstallDir -Force
}

# Remove from PATH environment variable
$envPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
$newEnvPath = ($envPath.Split(';') | Where-Object { $_ -ne $installDir }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newEnvPath, [EnvironmentVariableTarget]::Machine)

Write-Host "Uninstallation completed successfully!"
