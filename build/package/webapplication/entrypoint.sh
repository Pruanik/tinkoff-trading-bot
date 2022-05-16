#!/usr/bin/env bash

echo "[INFO] Start supervisord";
/usr/bin/supervisord -c /etc/supervisor/supervisord.conf;
echo "[INFO] Start supervisord finished";

echo "[INFO] Start WebApplication process";
./wait-for $DATABASE_HOST:$DATABASE_PORT -t 60 -- /usr/bin/supervisorctl start web:*
echo "[INFO] Start WebApplication process finished";

export SIGNAL_SENT="false";
#trap for processing system signal SIGTERM
trap 'signalCatch' SIGTERM
signalCatch() {
  pkill -SIGTERM -f supervisord
  export SIGNAL_SENT="true"
}

while true
do
  sleep 1
  if [ "$SIGNAL_SENT" == "true" ]; then
    ProcessesCount=$(pgrep webapplication | wc -l)
    if [ "$ProcessesCount" -gt 0 ];
    then
      # wait for some background processes which are still active
      sleep 10
    fi
    exit 0
  fi
done
