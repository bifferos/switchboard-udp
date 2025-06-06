Testing/Development
===================

$ make run

Will build and run a local copy with example config
from config.json.  That should give you a demo.
Tweak the values in config.json in case you have port
conflicts.


Installation
============

Only suported on Linux systemd.

$ make install

Will setup the system with some default locations. 
Install files manually if you want to change default locations.
The locations for the various directories are hard-coded
into the binary.  You can override them with the config file.
Anything missing from the config file will fall back 
to the defaults in main.go.

