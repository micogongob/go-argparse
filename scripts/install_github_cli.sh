#!/bin/bash

wget https://github.com/cli/cli/releases/download/v2.15.0/gh_2.15.0_linux_amd64.rpm
sudo rpm -i gh_2.15.0_linux_amd64.rpm
gh --version
