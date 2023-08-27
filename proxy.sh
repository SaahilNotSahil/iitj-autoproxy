#!/bin/bash

AUTOPROXY_DIR=/usr/bin

nohup $AUTOPROXY_DIR/autoproxyd & > $AUTOPROXY_DIR/autoproxyd.log
sleep 3
$AUTOPROXY_DIR/autoproxy login
sleep 5
sudo killall autoproxyd
