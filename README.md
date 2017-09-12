spacing-guild
===
Spacing Guild is a proxy to Memcached, written in Go.
It is intended to solve the issue of TCP connection flooding of Memcached from PHP.
Spacing Guild is currently in pre-release.

Requirements
------------
 * Install https://github.com/Masterminds/glide

Usage
-----
 * After performing a `go get` to retrieve the source, run `glide install`
 * Configurations are currently hard coded and will need to be updated if Memcached is not running locally at 11211
 * Run `go build main.go` to build then `./main` to run memproxy
