# UDP Server for state values

## Intro

Whilst one could use Home Assistant, Node Red or whatever for controlling
individual IOT devices it's sometimes quicker just to roll your own.  This
basic UDP server can be polled periodically by ESP-type devices to get a
true/false answer.  States are controlled via files on the server.
If a file exists, the value is 'true'.  If the file doesn't exist, the value is
'false'.  The text in the UDP request is the name of the file, the response
UDP request (to the sender port), is simply the text 'true' or 'false'.  
This is all completely insecure, of course!

## Requirements

You need Go for this, make sure it's installed, most package managers
will do that for you:

`$ sudo apt install golang-go`

## Testing

`$ make run`

Will build and run a local copy with example config
from config.json.  That should give you a demo.
Tweak the values in config.json in case you have port
conflicts.

`$ make test`

Will run a simple client to test the server.  If you forgot the
tcpdump command to listen to udp packets, try:

`$ make tcpdump`

## Installation

Only suported on Linux systemd.

`$ make install`

Will setup the system with some default locations. 

## Configuration

Post-install the default config should work fine.  The server will 
listen on port 6060, and check for existence of files in 
`/var/lib/switchboard` If you are unhappy with these defaults, 
copy config.json into the directory `/etc/switchboard-udp/config.json`
and edit the values to suit.

You will need to restart the service after changing config.json

`$ make stop`

`$ make start`
