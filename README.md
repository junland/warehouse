# Warehouse [![Build Status](https://travis-ci.org/junland/warehouse.svg?branch=master)](https://travis-ci.org/junland/warehouse) [![GolangCI](https://golangci.com/badges/github.com/junland/warehouse.svg)](https://golangci.com)

File and binary distribution service for people.

Warehouse tries to be a simple file server targeting distrbution of assets (CSS, Javascript, Pictures, etc.) and Unix / Linux operating system binary packages (RPM, DEB, IPK etc.). This simpliciaty enables developers or operations to deploy a simple file server over there existing pacakge or assets directory which will provide a clean interface to pull down assets or packages to devices. 

## Features

* Brandable file browser with custom JS, HTML, and CSS.

* Server binary under 10MiB.

* 

* Simple API to track downloads of files / packages. (Coming Soon!)

## Building

Binaries require Go (1.9 or newer) and Git. To build follow the instructions below:

1. Download the source and place it in your `GOPATH`.
2. Change directories into the source.
3. Run `go build`

You can also use the `Makefile` included in the source to build it.

## Getting started

To start serveing files from your current working directory, execute:

```
warehouse
```

Now you can look thru your directory under `http://localhost:8080/assets`.

To configure the server further issue this command:

```
warehouse --help
```

From there you can look thru the options to configure more directories.

## Built With

`github.com/justinas/alice` - Simple middleware chaining library.

`github.com/sirupsen/logrus` -  Structured, pluggable logging for Go.

`github.com/spf13/pflag` - Drop in replacement for the `flag` package.

`github.com/julienschmidt/httprouter` - A high performance HTTP request router that scales well.

## Versioning

I use [SemVer 2.0.0](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/junland/pak-mule/tags).

## Authors

* **John Unland** - *Initial work* - [junland](https://github.com/junland)

See also the list of [contributors](https://github.com/junland/warehouse/contributors) who participated in this project

## License

Code is licensed under GPLv2 which can be viewed in the `LICENSE` file.

_Please let me know through the issues tracker if you have any questions._

## TODO / Notes

* Not production ready.

* API integration.

* Write tests.

* Check issues list for more information.

