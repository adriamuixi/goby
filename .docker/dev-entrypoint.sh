#!/bin/bash
set -e
echo "WEB - Start downloading package -fresh- GOLANG"
go get -v github.com/pilu/fresh
echo "WEB - Finish downloading package -fresh- GOLANG"
echo "WEB - Starting fresh command"
fresh
echo "WEB - Started fresh command"
