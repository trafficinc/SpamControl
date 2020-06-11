# SpamControl
SpamControl will check for bad emails via a blacklist, Go app.

## Start

Change path to your SOCK directory: `const SOCK = "/path/to/spamcontrol/spam.sock"` in main.go & sockettest.php
 

## to build
$ go build


## test
$ go test -v

## Usage
1.) Open two consoles in project directory

2.) Run for Unix Socket: $ socat - UNIX-CONNECT:spam.sock in 1st console

3.) $ go run main.go in 2nd console

4.) Send test data: {"action":"url","value":"demo1@gmail.com"} 

Or you can run the php script to use with a scripting language to use the service. 

{make sure spam.sock file permissions are 0770 / $ chmod 0770 spam.sock}
