# TriggerShell

**TriggerShell** is a folder monitoring tool that automatically triggers shell scripts when new files are created in a directory. The script to execute is determined by the file extension.

## Features
- Monitors a directory for new files
- Maps file extensions to specific shell scripts
- Executes scripts automatically when matching files are detected

## Installation

Short Installation: Clone the repository, Navigate into the directory and build

   ```bash
   git clone https://github.com/DarkAxi0m/TriggerShell.git
   cd TriggerShell
   make
   sudo make install
   ```

Setup a new trigger
   ```bash
   cd TriggerShell
   make setup
   
   #Enter folder to watch (default: ~/Downloads): 
   #Enter folder for scripts (default: ./scripts): ~/scripts
   #TriggerShell.desktop file created and permissions set in /home/chris/.config/autostart with watch_dir=/home/chris/Downloads and scripts_dir=/home/chris/scripts
   #Do you want to start TriggerShell now? (y/n): n
   #TriggerShell setup complete. You can start it manually later.

   ```
Long Installation: (apt based) Install required tools, Clone the repository, Navigate into the directory and build    
Not including GO, follow here https://go.dev/doc/install

   ```bash
   apt update
   apt install git make
   #INSTALL GO
   git clone https://github.com/DarkAxi0m/TriggerShell.git
   cd TriggerShell
   make

   #Building TriggerShell ...
   #go build -o TriggerShell 
   #go: downloading github.com/fsnotify/fsnotify v1.7.0
   #go: downloading golang.org/x/sys v0.4.0

   sudo make install

   #Installing TriggerShell  to /usr/local/bin...
   #install -m 0755 TriggerShell  /usr/local/bin/TriggerShell

   TriggerShell
   #Usage: TriggerShell <target-directory> <script-directory>

   ```


## Usage
Run TriggerShell by specifying the target directory to monitor and the directory containing shell scripts:


   ```bash
   ./TriggerShell <target-directory> <script-directory>
   ```


* `target-directory`: The directory to be monitored for new files.
* `script-directory`: The directory containing shell scripts, which should be named with the file extension they will trigger (e.g., example.sh will trigger when a .example file is created).

## Examples
   
   ```bash
   ./TriggerShell ~/Downloads ./scripts
   ```

For example, in a scenario where you are watching your download folder, you could use the following scripts (found in the example folder of the repo):

- `appimage.sh`: Asks for confirmation before proceeding, then copies the file to `~/Applications`, sets the correct permissions, and creates a `.desktop` file.
- `torrent.sh`: When a `.torrent` file is downloaded, it uploads the file to a remote FTP server and then deletes it.

## Current bugs and to-dos
In no order

* Watch the scripts folder for changes
* ~~Stop the duplicate triggering with Chrome downloads~~
* Create Releases
* Create Deb install


## Contributing
Feel free to submit issues or pull requests.
