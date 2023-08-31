#!/bin/bash

AUTOPROXY_DIR=/usr/bin

# Start the daemon and redirect output to a log file
nohup $AUTOPROXY_DIR/autoproxyd > $HOME/autoproxyd.log &

# Get the PID of the last background command
APP_PID=$!

# Wait a bit for the daemon to start
sleep 3

# Execute the login command
if exec $AUTOPROXY_DIR/autoproxy login; then
    echo "Successfully logged in."
else
    echo "Failed to log in."
    exit 1
fi

# Wait some more time
sleep 5

# Kill the daemon
if kill $APP_PID; then
    echo "Successfully terminated the daemon."
else
    echo "Failed to terminate the daemon."
    exit 1
fi
