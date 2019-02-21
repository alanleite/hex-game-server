# Project Specs

Principal specs definitions on this project.

## Development

* Coverage minimal allowed is 99%
* Pull request reviews are obligated
* EditorConfig plugin is required on this project
* All new entry written on .env must be a example on .env.template.env preferably with development value on it
* Prefear commit small groups of files and separete with per scope change

## Documentation

* All documentation must have been writen and updated
* Development time must to take into account documentation time
* Swagger will be usade for REST documentation
* Socket documentation will be writen with md because swagger doesn't have a cool sockets plugin

## Golang

* Golang 1.10.7 will be used as minimal aceptable golang version for dev
* Solid will be used as main designe pattern
* [Golang Standards Project Layout](https://github.com/golang-standards/project-layout) will used as folders organizations
* [Dep](https://golang.github.io/dep) will be used as package manager instead of native goland package manger
* [Gorilla WebSocket](https://github.com/gorilla/websocket) will be used as socket framework
* [Realize Task Runner](https://github.com/oxequa/realize) will be used as task runner