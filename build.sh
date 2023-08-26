#!/bin/bash

echo "Building the autoproxy daemon..."

cd cli && go build -o ../bin/autoproxyd && cd ..

echo "Building the autoproxy cli..."

cd cli && go build -o ../bin/autoproxy && cd ..

echo "Build completed successfully!"

echo "Cleaning any previous installation..."

sudo rm -f /usr/local/bin/autoproxyd
sudo rm -f /usr/local/bin/autoproxy

echo "Installing now..."

sudo cp bin/autoproxyd /usr/local/bin/autoproxyd
sudo cp bin/autoproxy /usr/local/bin/autoproxy

echo "Creating config file..."

cp base_config.json $HOME/.autoproxy.config

echo "Config file copied to $HOME/.autoproxy.config"

echo "Installation completed successfully!"
