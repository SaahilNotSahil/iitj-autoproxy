#!/bin/bash

AUTOPROXY_DIR=/path/to/iitj-autoproxy

nohup $AUTOPROXY_DIR/autoproxyd & > $AUTOPROXY_DIR/autoproxyd.log
sleep 3
$AUTOPROXY_DIR/autoproxy-cli login
sleep 5
sudo killall autoproxyd
