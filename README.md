# Enumall
[![Go](https://github.com/tomaspavlic/enumall/actions/workflows/go.yml/badge.svg)](https://github.com/tomaspavlic/enumall/actions/workflows/go.yml)

Enumall is a tool to automate the creation of all const values for given type (enum). 

## Installation

`enumall` is installable command line application.

```
go install github.com/tomaspavlic/enumall
```

## Usage

Add Go's code generator comment to use `enumall`.

```golang
//go:generate go run github.com/tomaspavlic/enumall@latest -type=Season

type Season uint8

const (
    Spring Season = 1 << iota
    Summer
    Autumn
    Winter
)
```

Run code generator inside your module.
```
go generate ./...
```

Generated code is named `{$typeName}_all.go`. Variable contaings all const values is `All{$typeName}`

```golang
// Code generated by "enumall -type=Season"; DO NOT EDIT.

package main

var AllSeason = []Season{
	Spring,
	Summer,
	Autumn,
	Winter,
}

```