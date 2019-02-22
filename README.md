# Hex Game Server

This a Hex Game Server made in Golang

## See

* [Manifest](./manifest) folder
* [Project Specs](./manifest/project-specs.md)
* [Sockets Doc](./docs/socket/README.md)

## How Dev/

How start dev on this project

### Dependencies

* [Golang](https://golang.org) 1.10 or greater
* [Dep](https://golang.github.io/dep)
* [Realize Task Runner](https://github.com/oxequa/realize)

### Installing

`$ make packages`

### Running

Run the commands bellow

#### Server

This command will run the server on localhost:8080

`$ make dev`

#### Test Client

This command will connect at socket server and start to send and receive test connection data.

`$ make echo`

Will make an output like this it the server is on too:

```
2:47:52][HEX-GAME-CLIENT] : Watching 4 file/s 9 folder/s
[22:47:52][HEX-GAME-CLIENT] : Install started
[22:47:52][HEX-GAME-CLIENT] : Install completed in 0.087 s
[22:47:52][HEX-GAME-CLIENT] : Running..
[22:47:52][HEX-GAME-CLIENT] : connecting to ws://localhost:8080/echo
[22:47:53][HEX-GAME-CLIENT] : recv: 2019-02-21 22:47:53.946507627 -0300 -03 m=+1.001971906
[22:47:54][HEX-GAME-CLIENT] : recv: 2019-02-21 22:47:54.946425434 -0300 -03 m=+2.001889731
```

## DevOps

ToDO
