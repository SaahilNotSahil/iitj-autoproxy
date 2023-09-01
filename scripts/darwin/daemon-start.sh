#!/bin/bash

# This script is used to start the daemon

# Check if script is running as root
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 
   exit 1
fi

# Get a non-root username
TARGET_USER=$(who | awk '{if ($1 != "root") print $1; exit;}')

# Run the daemon as the target user
su - $TARGET_USER -c /opt/homebrew/bin/autoproxyd > /dev/null 2>&1 &

# Get the PID of the last background command
APP_PID=$!

# Write the PID to a file
echo $APP_PID > /var/run/autoproxyd.pid

# Disown the PID
disown $APP_PID
