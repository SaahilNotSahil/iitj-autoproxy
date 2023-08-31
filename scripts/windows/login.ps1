# Set the directory path
$AutoProxyDir = "C:\Path\To\AutoProxy"

# Start the daemon and redirect output to a log file
Start-Process -FilePath "$AutoProxyDir\autoproxyd.exe" -RedirectStandardOutput "$env:USERPROFILE\autoproxyd.log" -NoNewWindow

# Wait a bit for the daemon to start
Start-Sleep -Seconds 3

# Execute the login command
try {
    Start-Process -FilePath "$AutoProxyDir\autoproxy.exe" -ArgumentList "login" -Wait
    Write-Host "Successfully logged in."
}
catch {
    Write-Host "Failed to log in."
    exit 1
}

# Wait some more time
Start-Sleep -Seconds 5

# Kill the daemon
try {
    Stop-Process -Name "autoproxyd"
    Write-Host "Successfully terminated the daemon."
}
catch {
    Write-Host "Failed to terminate the daemon."
    exit 1
}
