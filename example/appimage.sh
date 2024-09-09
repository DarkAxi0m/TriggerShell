#!/bin/bash

if [ -z "$1" ]; then
    echo "No AppImage file provided!"
    exit 1
fi

APPIMAGE_FILE="$1"
APP_NAME=$(basename "$APPIMAGE_FILE" .AppImage)
APPLICATIONS_DIR=~/Applications
DESKTOP_FILE=~/Desktop/$APP_NAME.desktop

zenity --question --text="Do you want to install the AppImage for $APP_NAME?" --width=300
if [ $? -ne 0 ]; then
    echo "Installation canceled."
    exit 1
fi

APP_NAME=$(zenity --entry --text="Enter the application name" --entry-text="$APP_NAME" --width=300)
if [ -z "$APP_NAME" ]; then
    echo "No application name provided. Installation canceled."
    exit 1
fi

mkdir -p "$APPLICATIONS_DIR"
mv "$APPIMAGE_FILE" "$APPLICATIONS_DIR/"
chmod +x "$APPLICATIONS_DIR/$APP_NAME.AppImage"

cat <<EOF > "$DESKTOP_FILE"
[Desktop Entry]
Name=$APP_NAME
Exec=$APPLICATIONS_DIR/$APP_NAME.AppImage
Icon=$APPLICATIONS_DIR/$APP_NAME
Type=Application
Terminal=false
EOF

chmod +x "$DESKTOP_FILE"

zenity --info --text="AppImage installed successfully!\n\nApp: $APP_NAME\nLocation: $APPLICATIONS_DIR" --width=300
echo "AppImage moved to $APPLICATIONS_DIR and .desktop file created at $DESKTOP_FILE"

