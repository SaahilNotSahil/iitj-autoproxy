#!/bin/bash

# Install autoproxy daemon and CLI
echo "Installing now..."

sudo cp "bin/autoproxyd" "/usr/bin/autoproxyd"
sudo cp "bin/autoproxy" "/usr/bin/autoproxy"

echo "Creating example config directory..."

sudo mkdir -p "/etc/iitj-autoproxy"

echo "Creating example config file..."

sudo cp "autoproxy.config" "/etc/iitj-autoproxy/autoproxy.config"

echo "Config file copied to /etc/iitj-autoproxy/autoproxy.config"

echo "Copying helper scripts..."

sudo cp "scripts/darwin/daemon-start.sh" "/usr/bin/autoproxyd-start"
sudo cp "scripts/darwin/daemon-stop.sh" "/usr/bin/autoproxyd-stop"
sudo cp "scripts/darwin/login.sh" "/usr/bin/autoproxyd-login"

echo "Installation completed successfully!"
