#!/bin/bash
if pgrep -f web > /dev/null
then
    echo "Application is running"
    exit 0
else
    echo "Application is not running"
    exit 1
fi