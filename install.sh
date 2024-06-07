
# Check if script is run as root
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root. Restarting as root..."
    sudo "$0" "$@"
    exit $?
fi

# Define directories
binDir="bin"
scriptsDir="scripts"
serviceFilesDir="service_files"
installDir="/usr/bin"
serviceDir="/etc/systemd/system"
configInstallDir="/etc/iitj-autoproxy"
readmeDir="/usr/share/doc/iitj-autoproxy"
licenseDir="/usr/share/licenses/iitj-autoproxy"

# Create installation directories if they don't exist
mkdir -p "$configInstallDir"
mkdir -p "$readmeDir"
mkdir -p "$licenseDir"

# Copy binaries
cp "$binDir/autoproxy" "$installDir"
cp "$binDir/autoproxyd" "$installDir"
cp "$scriptsDir/linux/daemon-start.sh" "$installDir/autoproxyd-start"
cp "$scriptsDir/linux/daemon-stop.sh" "$installDir/autoproxyd-stop"
cp "$scriptsDir/linux/login.sh" "$installDir/autoproxylogin"
cp "$serviceFilesDir/linux/autoproxyd@.service" "$serviceDir/autoproxyd@.service"
cp "autoproxy.config" "$configInstallDir"
cp "README.md" "$readmeDir"
cp "LICENSE" "$licenseDir"

# Manage permissions
chmod +x "$installDir/autoproxyd-start"
chmod +x "$installDir/autoproxyd-stop"
chmod +x "$installDir/autoproxylogin"

# Starting service
systemctl daemon-reload
systemctl enable --now autoproxyd@$(whoami).service

echo "Installation of IITJ Autoproxy completed successfully!"
