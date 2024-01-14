#!/bin/bash

# shellcheck disable=SC2181
# shellcheck disable=SC2006

function start_app() {
    pid=`ps -ef | grep WePanel | grep -v grep | awk '{print $2}'`
    if kill -0 ${pid} > /dev/null 2>&1
    then
        echo "[error] app is running, pid=${pid}"
        return 1
    fi
    echo "[info] start_app begins..."
    go build -o WePanel backend/main.go
    nohup /home/WePanel/WePanel >WePanel.log&
    if [ $? -ne 0 ];then
        echo "[error] start app failed!"
        return 1
    fi
    sleep 1
    pid=`ps -ef | grep WePanel | grep -v grep | awk '{print $2}'`
    echo "[info] app start at pid: [${pid}]"
}

function stop_app() {
    pid=`ps -ef | grep WePanel | grep -v grep | awk '{print $2}'`
    echo "[info] stop_app begins..."
    if ! ( kill -0 ${pid} )
    then
        echo "[error] app is not running, please start first."
        return 1
    fi
    echo "[info] pid is [${pid}], begin kill it..."
    restart_pid=`ps -ef | grep pv_restart | grep -v grep | awk '{print $2}'`
    kill -9 "${pid}"
    if [ $? != 0 ]; then
      echo "[error] pid is [${pid}], kill it failed, try to kill it manually!"
      return 1
    fi

    echo "[info] app is killed successfully."
}

function restart_app() {
    pid=`ps -ef | grep WePanel | grep -v grep | awk '{print $2}'`
    if ! ( kill -0 ${pid} )
    then
        echo "[info] app is not running, then start app."
        start_app
        if [ $? != 0 ]; then
            return 1
        fi
    else
        echo "[info] app is running, stop app first."
        stop_app
        if [ $? != 0 ]; then
            return 1
        fi
        start_app
        if [ $? != 0 ]; then
            return 1
        fi
    fi
    echo "[info] restart app successfully."
}

function main() {
    typeset cmd_type=$1
    case "${cmd_type}" in
        "stop")
          stop_app
          if [ $? != 0 ]; then
              echo "[error] stop_app failed!"
              return 1
          fi
        ;;
        "start")
          start_app
          if [ $? != 0 ]; then
              echo "[error] start_app failed!"
              return 1
          fi
        ;;
        "restart")
          restart_app
          if [ $? != 0 ]; then
              echo "[error] restart_app failed!"
              return 1
          fi
        ;;
        *)
          echo "[error] [${cmd_type}] is not in [start, stop, restart]!"
          return 1
        ;;
    esac
    echo "[info] exec [${cmd_type}] successfully."
}

main "$@"