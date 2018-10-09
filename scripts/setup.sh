#!/bin/bash

setupMac() {

}

setupLinux() {

}

if [ "$(uname -s)" == "Darwin" ] then
    setupMac
else
    setupLinux
fi