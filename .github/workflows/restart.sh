#!/bin/bash
# this is a helper shell for github action in 'go.yml'
cd ~/auto-deploy/backend/
kill $(cat pid)
nohup ./cofmgr > cofmgr.log 2>&1 &
jobs -p > pid
echo Finished