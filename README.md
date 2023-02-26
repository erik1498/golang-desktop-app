# golang-desktop-app
Make WebApp Like Desktop App With Golang


UBUNTU
  - apt install libgtk-3-dev libwebkit2gtk-4.0-dev
  - apt install libgtk-3-0 libwebkit2gtk-4.0-37


Make file in this directory "golang-desktop-app.desktop"

Write

[Desktop Entry]
Encoding=UTF-8
Version=1.0
Type=Application
Terminal=false
Exec=/home/erickhene/go/src/golang-desktop-app
Icon=/home/erickhene/go/src/golang-desktop-app/favicon.ico
Name=phpMyAdmin

Save and Close

run "sudo cp golang-desktop-app.desktop ~/local/share/applications/
