#!/bin/bash
# this is a helper shell for github action in 'go.yml'
cd ~/auto-deploy/backend/
kill $(cat pid)
echo Finished