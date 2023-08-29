#!/bin/bash

set -e # Exit on any error

# Function to build a Go application
build_go_app() {
    local dir="$1"
    local output="$2"

    echo "Building $output..."
    (
        cd "$dir" || exit 1
        go build -o "$output" || exit 1
    )
}

# Function to safely remove a file
safe_remove() {
    local file="$1"
    [ -f "$file" ] && sudo rm -f "$file"
}

# Clean up previous installation
echo "Cleaning any previous installation..."

safe_remove "/usr/bin/autoproxyd"
safe_remove "/usr/bin/autoproxy"
sudo rm -rf "/etc/iitj-autoproxy"

# Build autoproxy daemon and CLI
echo "Building the autoproxy daemon and CLI..."

build_go_app "daemon" "../bin/autoproxyd"
build_go_app "cli" "../bin/autoproxy"

echo "Build completed successfully!"

# Install autoproxy daemon and CLI
echo "Installing now..."

sudo cp "bin/autoproxyd" "/usr/bin/autoproxyd"
sudo cp "bin/autoproxy" "/usr/bin/autoproxy"

echo "Creating example config directory..."

sudo mkdir -p "/etc/iitj-autoproxy"

echo "Creating example config file..."

sudo cp "autoproxy.config" "/etc/iitj-autoproxy/autoproxy.config"

echo "Config file copied to /etc/iitj-autoproxy/autoproxy.config"

echo "Installation completed successfully!"
