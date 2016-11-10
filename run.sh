#!/bin/bash
if [ ! -e "token" ]
then
    echo "File 'token' does not exist"
    exit
fi
xargs go run bot.go -token < token
