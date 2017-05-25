#! /bin/bash
echo 'build app'
go build
echo 'package app for windows'
bee pack -be GOOS=windows -exp=.:upload:logs:bee.json:package_linux.sh:package_windows:README.md:.git
