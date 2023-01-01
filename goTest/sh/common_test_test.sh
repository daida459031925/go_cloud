#!/bin/bash

#1.检测是否安装
#ls /usr/bin |grep expect
#如果不存在，则进行安装
#2.安装
#sudo apt-get install expect

/usr/bin/expect << EOF
# 设置超时时间
set timeout 3
set host "192.168.2.106"
set name "dell"
set pwd  "87793891"
# fork一个子进程执行ssh
spawn ssh $name@$host

expect {
  "*(yes/no/*"
    { send "yes\r" }
  "*password*"
    { send "$pwd\r" }
  "*#"
    { send "$pwd" }
}
# 如果时第一次进入那么就设置yes
#expect "*(yes/no/*" {
#  send "yes\r"
#}
## 设置密码
#expect "*password*" {
#  send "$pwd\r"
#}
## 打印自己的密码
#expect "*#" {
#  send "$pwd"
#}

expect eof

send "cd /home\r"

# 允许用户进行交互
interact
EOF

#!/bin/bash

# 设置超时时间
set timeout 3
set host "192.168.2.106"
set name "dell"
set pwd  "87793891"
# fork一个子进程执行ssh
spawn ssh $name@$host

expect {
  "*(yes/no/*"
    { send "yes\r" }
  "*password*"
    { send "$pwd\r" }
  "*#"
    { send "$pwd" }
}
# 如果时第一次进入那么就设置yes
#expect "*(yes/no/*" {
#  send "yes\r"
#}
## 设置密码
#expect "*password*" {
#  send "$pwd\r"
#}
## 打印自己的密码
#expect "*#" {
#  send "$pwd"
#}

expect eof

send "cd /home\r"

# 允许用户进行交互
interact
