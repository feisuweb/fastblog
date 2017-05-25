#! /bin/bash
echo 'build app'
go build
echo 'package app for linux'
bee pack -be GOOS=linux -exp=.:upload:logs:bee.json:package_linux.sh:package_windows:README.md:.git
