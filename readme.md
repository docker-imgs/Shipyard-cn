# Shipyard中文版本
Composable Docker Management

[![Build Status](https://travis-ci.org/shipyard/shipyard.svg?branch=master)](https://travis-ci.org/shipyard/shipyard)

```
curl -s http://panli.mu.gg/docker/deploy | bash -s
```

For full options:

```
curl -s http://panli.mu.gg/docker/deploy | bash -s -- -h
```

###  [领说明](http://panli.mu.gg/2016/02/18/docker/docker%E7%AE%A1%E7%90%86%E5%B9%B3%E5%8F%B0-shipyard-%E4%B8%AD%E6%96%87%E7%89%88%E9%83%A8%E7%BD%B2/)




# Documentation
Full docs are available at http://shipyard-project.com

# Components
There are three components to Shipyard:

## Controller
The Shipyard controller talks to a RethinkDB instance for data storage (user accounts, engine addresses, events, etc).  It also serves the API and web interface (see below).  The controller uses Citadel to communicate to each host and handle cluster events.

## API
Everything in Shipyard is built around the Shipyard API.  It enables actions such as starting, stopping and inspecting containers, adding and removing engines and more.  It is a very simple RESTful JSON based API.

## UI
The Shipyard UI is a web interface to the Shipyard cluster.  It uses the Shipyard API for all interaction.  It is an AngularJS app that is served via the Controller.

# Contributing

## Controller
To get a development environment you will need:

* Go 1.4+
* Node.js: (npm for bower to build the Angular frontend)

Run the following:

* install [Godep](https://github.com/tools/godep): `go get github.com/tools/godep`
* run `npm install -g bower` to install bower
* run `make build` to build the binary
* run `make media` to build the media
* run `./controller -h` for options

# License
Shipyard is licensed under the Apache License, Version 2.0. See LICENSE for full license text.
