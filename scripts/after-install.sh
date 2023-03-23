#!/bin/bash
echo "updating file permissions"
chown -R ubuntu:ubuntu /home/ubuntu/crudapp/
chmod +x /home/ubuntu/crudapp/main
