# gopasswd
go program to generate a password for use in /etc/shadow

# how to use
right now only env var input is supported, might add config file support at some point

## Manually
- run `go get` to download the dependencies
- set the variable `PASSWD` (example: `export PASSWD="password123"`)
- run `go run main.go` to get the output
- copy the output
- change the password field (the second one) in `/etc/shadow` to what you copied
- log out
- log back in with your new password
- enjoy!

## Automatically
- run `go get` to download the dependencies
- set the environment variable `PASSWD` (example: `export PASSWD="password123"`)
- set the environment variable `USR` (example: `export USR="user123"`)
- set the environment variable `REPLACE_IN_SHADOW` (example: `export REPLACE_IN_SHADOW="justneedstonotbeempty"`)
- run `go run main.go`
- log out
- fill your heart with hope
- log back in with your new password
- enjoy(?)
