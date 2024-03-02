# tri [![Go](https://github.com/shina-blueprint/tri/actions/workflows/go.yml/badge.svg)](https://github.com/shina-blueprint/tri/actions/workflows/go.yml)

**Ternary operator-like functions in Go**

## Overview

This Go module provides four functions that serve as alternatives to the ternary operator.

Each function evaluates a condition and returns one of two values or executes an argument-less function (for deferred execution).

## Installation

```
go get github.com/shina-blueprint/tri
```

## Usage

```go
package main

import "github.com/shina-blueprint/tri"

func main() {
    // a = "true value"
    a := tri.If(true, "true value", "false value")
    // b = "false value"
    b := tri.If(false, "true value", "false value")

    // c = "true func"
    c := tri.IfExec(true, func() string { return "true func" }, func() string { return "false func" })
    // d = "false func"
    d := tri.IfExec(false, func() string { return "true func" }, func() string { return "false func" })

    // e = "true func"
    e := tri.IfExecOrVal(true, func() string { return "true func" }, "false value")
    // f = "false value"
    f := tri.IfExecOrVal(false, func() string { return "true func" }, "false value")

    // g = "true value"
    g := tri.IfValOrExec(true, "true value", func() string { return "false func" })
    // h = "false func"
    h := tri.IfValOrExec(false, "true value", func() string { return "false func" })
}
```
