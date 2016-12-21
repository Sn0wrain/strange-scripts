## Preparing Go env for a brand new Debian
   趁着双11特价购入一年阿里云ECS,花费一小时搭建好了工作环境,现将配置过程呈上,方便之后参考.

- 创建用户
    由于root权限太高,容易造成严重后果,故最好使用非root用户进行操作:
```bash
useradd myuser
passwd myuser
```

- 修改用户sh
    Debian创建的用户默认采用dash,在实用性方面不如bash或zsh等,故需切换dash:
```bash
vim /etc/passwd
将 myuser:x:1000:1000:myuser,,,:/home/myuser:/bin/sh
改为 myuser:x:1000:1000:myuser,,,:/home/myuser:/bin/bash
```

- Sudo
    Debian默认环境是没有sudo命令的,因此也无法给普通账户以管理员权限,故需先安装sudo,再配置账户权限：
```bash
aptitude install sudo

vim /etc/sudoers

root    ALL=(ALL:ALL) ALL
myuser  ALL=(ALL:ALL) ALL
```

- 工具
    工作时需要的一些小工具,如:vim, git, fcitx-googlepinyin等,采用apt-get或aptitude安装即可.

- Go
    Go语言配置环境需先安装Golang 1.4,安装好之后先配置Go_bootstrap环境:
```bash
sudo vim /etc/profile
加上
export GOROOT_BOOTSTRAP=/usr/local/go1.4
```
    之后即可安装高版本的Golang 1.7,完成安装后需将剩余环境配好:
```bash
sudo vim /etc/bash.bashrc
```
加上
```bash
export GOROOT=/usr/local/go
export GOARCH=amd64
export GOOS=linux
export GOBIN=$GOROOT/bin
export PATH=$GOBIN:$PATH
```

- Beanstalkd
    Beanstalkd是一个轻量级消息中间件,它最大特点是将自己定位为基于管道(tube)和任务(job)的工作队列(work-queue).
    Beanstalkd支持任务优先级(priority),延时(delay),超时重发(time-to-run)和预留(buried),能够很好的支持分布式的后台任务和定时任务处理。它的内部实现采用libevent,服务器-客户端之间用类似memcached 的轻量级通讯协议,具有有很高的性能.尽管是内存队列,beanstalkd提供了binlog机制,当重启beanstalkd时,当前任务状态能够从纪录的本地binlog中恢复.
```bash
sudo vim /etc/default/beanstalkd
```
加入
```bash
START=yes
```
之后即可采用命令控制Beanstalkd服务了.
```bash
sudo /etc/init.d/beanstalkd start
lsof -i:11300
sudo /etc/init.d/beanstalkd stop
```

- Supervisor
    Linux的后台进程运行有好几种方法，例如nohup，screen等，但是，如果是一个服务程序，要可靠地在后台运行，我们就需要把它做成daemon，最好还能监控进程状态，在意外结束时能自动重启。
    supervisor就是用Python开发的一套通用的进程管理程序，能将一个普通的命令行进程变为后台daemon，并监控进程状态，异常退出时能自动重启。
    安装后需要给自己的Application编写一个config文件,以便Supervisor管理,每个进程的配置文件都可以单独分拆,放在/etc/supervisor/conf.d/目录下,以.conf作为扩展名,例如,app.conf定义了一个gunicorn的进程：
```bash
[program:app]
command=/usr/bin/gunicorn -w 1 wsgiapp:application
directory=/srv/www
user=www-data
```
