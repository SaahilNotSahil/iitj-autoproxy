#!/bin/bash

AUTOPROXY_DIR=/usr/bin

nohup $AUTOPROXY_DIR/autoproxyd & > $HOME/autoproxyd.log
sleep 3
exec $AUTOPROXY_DIR/autoproxy login
sleep 5
sudo killall autoproxyd
