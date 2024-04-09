#!/bin/sh

socket_open() {
    local host="${1%:*}"
    local port="${1#*:}"
    printf "HEAD / HTTP/1.0\r\n\r\n" \
        | nc -w 3 -n "$host" "$port" \
        | grep "Proxy Authentication Required" > /dev/null
}

socket="$1"
interval="$2"
executable="$3"
args="$4"

while true; do
    if socket_open "$socket"; then
        echo "$socket is open"
        if ! pgrep "${executable##*/}" > /dev/null; then
            echo "Starting $executable $args"
            $executable $args &
        fi
    else
        echo "$socket is close"
        if pgrep "${executable##*/}" > /dev/null; then
            echo "Stopping $executable..."
            pkill "$executable"
        fi
    fi

    echo "Sleeping $interval seconds"
    sleep "$interval"
done
