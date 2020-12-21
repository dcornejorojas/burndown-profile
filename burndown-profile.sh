#! /bin/sh
# /etc/init.d/burndown-profile
#
USER="user"
GOPATH="/opt/go"
GOROOT="$GOPATH"
WORKDIR="$GOPATH/bin/"
# Carry out specific functions when asked to by the system
case "$1" in
  start)
    cd "$WORKDIR"
    echo "Start Golang script"
    export GOPATH="$GOPATH"
    export PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin"
    #/bin/su -m -l $USER -c "nohup ./myapp > $WORKDIR/myapp.log 2>&1 & echo $! > /path/to/apps/pids/myapp.pid"
    /bin/su -m -l $USER -c "nohup ./profile > /dev/null 2>&1 & echo $! > /opt/pids/profile.pid"
    ;;
  stop)
    echo "Stop Golang script"
    kill -9 `cat /opt/pids/profile.pid`
    ;;
  *)
    echo "Usage: /etc/init.d/burndown-profile {start|stop}"
    exit 1
    ;;
esac
exit 0