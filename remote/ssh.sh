#!/usr/bin/expect -f

# A ssh script for auto login to remote host.
# Author : Sn0wrain
# Mail   : xjpgogle@gmail.com

# Dependency : expect

set port xxxx 
set user xxx
set host xx.xx.xx.xx
set password xxxxx
set timeout -1

spawn ssh -p $port $user@$host
expect "*assword:*"
send "$password\r"
interact
expect eof
