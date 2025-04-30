#!/bin/bash

############
## ./watchdog.sh "/Users/wuh/study/shell/script/test.php" 2 /tmp/sh/test.log "123123213" restart
#############

# $0 表示脚本本身的名称
# $1 待执行脚本
# $2 进程数量
# $3 日志文件
# $4 unique id

script=$1
processes=$2
logfile=$3
unique_id=$4
restart=$5
BIN_PATH=$(which php)

# 检查脚本是否为空
if test -z "$script" 
then
    echo "usage: $0 <script> <processes> <logfile> <unique_id> [restart]"
    exit 1
fi

# 日志文件是否存在
echo $logfile
if test ! -f "${logfile}"
then
    # 创建日志文件
    mkdir -p "$(dirname "$logfile")" && touch "$logfile"
    echo "创建日志文件 ${logfile}"
fi
# 执行命令并获取结果
# printf "%s" "ps -ef | grep "$BIN_PATH $script" | grep -v -c grep"
count=$(ps -ef | grep "$BIN_PATH $script" | grep -v -c grep)
# pids1=$(ps -ef | grep "$BIN_PATH $script" | grep -v grep | awk '{print $2}')
pids=$(pgrep -f "$BIN_PATH $script")


#重启进程 
if [ -n "$restart" ] && [ "$restart" = "restart" ]
then
    for pid in $pids
    do 
        echo "停止进程: $pid"
        kill -9 $pid
    done
    
    i=1
    while(( $i <= $processes )) 
    do 
        # 启动进程
        echo "启动进程: $BIN_PATH $script $unique_id"
        nohup $BIN_PATH $script $unique_id >> $logfile 2>&1 &
        i=$((i + 1))
        sleep 1
    done
    echo "restart process: $processes"
    exit 0
fi

if [ $processes -gt $count ] #进程数大于当前进程数
then
    while(( $count < $processes )) 
    do
        # 启动进程
        echo "启动进程: $BIN_PATH $script $unique_id"
        nohup $BIN_PATH $script $unique_id >> $logfile 2>&1 &
        count=$((count + 1))
        sleep 1
    done
elif [ $count -gt $processes ]
then
    echo "当前进程数: $count, 需要停止进程数: `expr $count - $processes`"
    # 停止多余的进程
    echo $pids
    for pid in $pids
    do
        # 只保留前 $processes 个进程
        if [ $count -gt $processes ]
        then
            echo "停止进程: $pid"
            kill -9 $pid
            count=$((count - 1))
        fi
    done
else
    echo "当前进程数: $count, 需要启动进程数: $processes"
    exit 0
fi

