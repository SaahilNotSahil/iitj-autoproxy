#!/bin/bash

# This script is used to stop the daemon

APP_PID=$(cat /run/autoproxyd.pid)

kill $APP_PID
