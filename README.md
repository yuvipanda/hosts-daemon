# hosts-daemon
Allows using a /etc/hosts.d to keep hosts aliases

## Why? ##

So you can more easily manage `/etc/hosts` aliases for whatever
purposes without having to do it all with one big file.

You can put individual files in `/etc/hosts.d/` and they'll all
be concated together to form `/etc/hosts`.

## How? ##

The daemon just has a timer that ticks (every 5s by default), 
reading the contents of `/etc/hosts.d/` and writing it out
to `/etc/hosts`. That's it!
