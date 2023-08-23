# gomod-viz

A CLI tool that lets you visualize the dependency graph of your project.

## Requirements
This tool requires you to have [Go](https://go.dev/doc/install) and [graphviz](https://graphviz.org/download/) installed, in order to generate the svg file.
It currently only works on Unix-like (Mac, Linux) systems

## Installation and Usage
To use gomod-viz, clone this repository and from the gomod-viz directory, run
```
make build
```
This will create a binary and store it in your local bin directory (/usr/local/bin).
To run the program, move to the directory of one of your Go projects, and run
```
gomod-viz
```

A svg file should open in whatever default viewing program is set on your machine.
