## Description
This is a simple tool to write some thoughts on the go from a terminal in your mycorrhiza diary.

## Installation
1. Place binary executable somewhere on the PATH
2. Retrieve your auth token from mycorrhiza. It can be found using browser tools - Inspect - Network - Some wiki request - Cookies
3. Recommended: Add user-wide alias for simpler usage. 
    - Linux: `alias diary='diary https://example.wiki "somelongtoken"'`
    - Windows: TODO

## Usage
If you already set up aliases, just execute `diary Some thoughts, questions, notes, etc`.
Else execute full command: `diary https://example.wiki "somelongtoken" "root/diary/hypha/name" Some thoughts, questions, notes, etc`
It will find or create hypha with current date as a name and then add a new line with current time at the line start.