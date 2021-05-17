# TOP -- Knowing your server's workload

A simple web-based application, that allows to monitor the system it runs on by rendering the output of the `top` command into a web page.

## Introduction

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. Enjoy experimenting!

**Your help is required.**

Can you please check whether you can find any vulnerability and repair anything you notice.

## Prerequisites

What things you need to have in the first place:

* [Golang](https://golang.org/) (Version 1.12 or above)
* [Gorilla Toolkit - Multiplexer](https://www.gorillatoolkit.org/pkg/mux)

## Installation

Follow these steps:

* The folder in which you unpacked the web app is denoted as "$TODOPATH" in the following (something like "/path/to/swa__prakt5_top" for example).

* Enter the web application's directory

```CLI
cd $TODOPATH
```

* Build the application from the sources

```CLI
go build
```

* Start the server (you could also start the executable "swa__prakt5_top" directly)

```CLI
go run topServer.go
```

* Access the web application using your browser: [http://localhost:8181/](http://localhost:8181/)

## Built with

* [Golang](https://www.golang.org/) - The Go Programming Language

## Author

* **Luigi Lo Iacono**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
