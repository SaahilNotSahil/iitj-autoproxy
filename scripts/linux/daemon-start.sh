#!/bin/bash

# This script is used to start the daemon

TARGET_USER=$(who | awk '{if ($1 != "root") print $1; exit;}')

su - $TARGET_USER -c /usr/bin/autoproxyd > /dev/null 2>&1 &

APP_PID=$!

echo $APP_PID > /run/autoproxyd.pid

disown $APP_PID
