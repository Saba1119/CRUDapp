#!/bin/bash
echo "starting go application"
/home/ubuntu/crudapp/main > /dev/null 2> /dev/null < /dev/null &
/home/ubuntu/crudapp/imageupload > /dev/null 2> /dev/null < /dev/null &
