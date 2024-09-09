# TriggerShell

**TriggerShell** is a folder monitoring tool that automatically triggers shell scripts when new files are created in a directory. The script to execute is determined by the file extension.

## Features
- Monitors a directory for new files
- Maps file extensions to specific shell scripts
- Executes scripts automatically when matching files are detected

## Installation

Clone the repository, Navigate into the directory and build
   ```bash
   git clone https://github.com/DarkAxi0m/TriggerShell.git
   cd TriggerShell
   go build -o TriggerShell
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
   ./TriggerShell ~/Downloads ./scripts`
   ```

For example, in a scenario where you are watching your download folder, you could use the following scripts (found in the example folder of the repo):

- `appimage.sh`: Asks for confirmation before proceeding, then copies the file to `~/Applications`, sets the correct permissions, and creates a `.desktop` file.
- `torrent.sh`: When a `.torrent` file is downloaded, it uploads the file to a remote FTP server and then deletes it.



## Contributing
Feel free to submit issues or pull requests.
