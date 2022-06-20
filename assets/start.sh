#!/bin/bash
# this is a helper shell script for github action in 'go.yml'
cd ~/auto-deploy/backend/
nohup ./cofmgr >> cofmgr.log 2>&1 &
jobs -p > pid
echo Finished