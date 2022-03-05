# gopasswd
go program to generate a password for use in /etc/shadow

# how to use
(right now only env var input is supported)
- set the variable `PASSWD` and then run `go run main.go` to get the output
- copy the output
- change the password field in `/etc/shadow` to what you copied
- log back in with your new password
- enjoy!
