#!/bin/bash

### CURRENTLY ONLY FOR DEBIAN-BASED LINUX DISTRIBUTIONS

echo "Building the autoproxy daemon..."

cd cli && go build -o ../bin/autoproxyd && cd ..

echo "Building the autoproxy cli..."

cd cli && go build -o ../bin/autoproxy && cd ..

echo "Build completed successfully!"

echo "Cleaning any previous installation..."

sudo rm -f /usr/bin/autoproxyd
sudo rm -f /usr/bin/autoproxy

echo "Installing now..."

sudo cp bin/autoproxyd /usr/bin/autoproxyd
sudo cp bin/autoproxy /usr/bin/autoproxy

echo "Creating example config file..."

cp base_config.json /etc/iitj-autoproxy/autoproxy.config

echo "Config file copied to /etc/iitj-autoproxy/autoproxy.config"

echo "Installation completed successfully!"
