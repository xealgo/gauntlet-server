# gauntlet-server
Multiplayer game server written in Go. Checkout the client [here](https://github.com/xealgo/gauntlet).

## Notes
* This is very early on. I just began rewriting the server in Go as opposed to C++ and this will be my first complex project dealing with UDP so I'm sure there will be a lot of dumb mistakes.

## Setup
The follow are the general setup and build commands. You can optionally run the commands
found within the makefile directly in your terminal, but I prefer makefiles for...making things ;)

#### Deps
```bash
> go get -u ./...
```

#### Run  
```bash
> make run
```

#### Install
```bash
> make install
```

#### Test
```bash
> make test
```
