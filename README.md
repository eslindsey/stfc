# stfc

Go client for Star Trek: Fleet Command

## Goals

The goal of this project is to allow developers to build software to expand the
community around the Scopely video game Star Trek: Fleet Command. This library
exposes endpoints utilized by the game client itself to obtain information
available to the player it is logged in as. The library is quite obvious in its
activies, and as such should be utilized in a *responsible* fashion. The author
of this library does not endorse and will not assume responsibility for misuse
of it. 

Some of the projects this library could be useful for: Discord bots, stats
tracking websites, custom leaderboards.

## Authentication

Currently, direct login using Scopely ID is not supported. You must figure out
how to obtain the `adhoc_username` and `adhoc_password` that your game client
uses to establish a session every time it starts up, and provide those
credentials to the library.

## Example Use

```go
session, err := stfc.Login(AdhocUsername, AdhocPassword)
// handle errors
fmt.Printf("Version:    %s\n", session.LoginResponse.Version)
fmt.Printf("Session ID: %s\n", session.LoginResponse.InstanceSessionID)
fmt.Printf("Name:       %s\n", session.LoginResponse.InstanceAccount.Name)

sync, err := session.Sync(2)
// handle errors
fmt.Printf("You have %d types of resources.\n", len(sync.Resources))
```

For more, see [stfc_test.go](stfc_test.go).

## Disclaimer

This material is unofficial and is not endorsed by Scopely, Paramount, Digit,
or any other entity. This work is provided on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including,
without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT,
MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE. You are solely
responsible for determining the appropriateness of using this work and assume
any risks associated with your use of it.

