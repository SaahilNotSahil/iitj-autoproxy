#!/bin/bash

# Check if script is running as root
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 
   exit 1
fi

# Check if PID file exists
if [ ! -f /run/autoproxyd.pid ]; then
    echo "PID file not found. Is the daemon running?"
    exit 1
fi

# Read PID from file
APP_PID=$(cat /run/autoproxyd.pid)

# Check if the PID is actually running
if ! ps -p $APP_PID > /dev/null; then
    echo "No process with PID $APP_PID found."
    exit 1
fi

# Kill the process
if kill $APP_PID; then
    echo "Successfully terminated process with PID $APP_PID."
    rm /run/autoproxyd.pid
else
    echo "Failed to terminate process with PID $APP_PID."
    exit 1
fi
