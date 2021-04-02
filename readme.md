# rick

`rick` provides a command line interface for controlling Spotify without having to open the window.

## setup
In the current version, environment variables `SPOTIFY_ID` and `SPOTIFY_SECRET` must be set, can be found in the spotify developer section.
Additionally, the project must be setup with `http://localhost:8080/callback` as a callback URL.

This is only a temporary measure until proper authentication is implemented, should probably use some safe credential storage 

## TODO
* improved authentication
* handle config files with viper?

