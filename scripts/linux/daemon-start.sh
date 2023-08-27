#!/bin/bash

# This script is used to start the daemon

/usr/bin/autoproxyd > /dev/null 2>&1 &

APP_PID=$!

echo $APP_PID > /run/autoproxyd.pid

disown $APP_PID
