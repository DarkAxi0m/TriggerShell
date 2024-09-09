#!/bin/bash

if [ -z "$1" ]; then
    echo "No torrent file provided!"
    exit 1
fi

FILENAME="$1"

#move uploaded torrents to ehre
TORRENT=~/Downloads/Torrent
SERVER=remote.torrent-server.com         

hash=$(md5sum "${FILENAME}"|cut -d' ' -f1)
ext=${FILENAME##*.}
mv -v "$FILENAME" "$TORRENT/${hash}.${ext}"

cd $TORRENT
echo ===Detected $FILENAME ===
echo ${hash}.${ext}

ftp <<EOF
open ${SERVER}   
cd watch 
put ${hash}.${ext}
exit
EOF

notify-send --expire-time=1000 "Torrent $FILENAME" 
