# rick

`rick` provides a command line interface for controlling Spotify without having to open the window.

## build from source
Requires at least Go 1.16

In order to build from source, clone down the repository and run `go install` from the root directory.
Make sure the install directory is on the path in order to call directly from the command line.


## setup
In the current version, environment variables `SPOTIFY_ID` and `SPOTIFY_SECRET` must be set, can be found in the spotify developer section.
Additionally, the project must be setup with `http://localhost:8080/callback` as a callback URL.

This is only a temporary measure until proper authentication is implemented, should probably use some safe credential storage 

## TODO
* improved authentication
* handle config files with viper?

