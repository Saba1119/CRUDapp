#!/bin/bash
echo "stopping go application"
ps -ef | grep main| grep -v grep | awk '{print $2}' | xargs kill
ps -ef | grep imageupload| grep -v grep | awk '{print $2}' | xargs kill
