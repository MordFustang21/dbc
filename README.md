# Design by Contract (dbc)

[![GoDoc](https://godoc.org/github.com/MordFustang21/dbc?status.svg)](https://godoc.org/github.com/MordFustang21/dbc)

A lightweight, zero-dependency Go implementation of the Design by Contract (DbC) programming paradigm.

## Overview

Design by Contract is a software design approach that focuses on clearly defining contracts for software components. The `dbc` package provides tools to express and enforce these contracts through:

- **Preconditions** (requirements a function caller must satisfy)
- **Postconditions** (guarantees a function makes about its results)
- **Invariants** (conditions that must remain true throughout execution)

## Features

- **Zero production overhead**: All checks except `Check()` become no-ops in production builds
- **Clear failure messages**: Panics include function name, file, line number, and formatted message
- **Simple API**: Just four functions to learn - `Require()`, `Ensure()`, `Invariant()` and `Check()`
- **Format string support**: Error messages use `fmt.Sprintf`-style formatting

## Installation

```bash
go get github.com/MordFustang21/dbc
```

## Usage

```go
import "github.com/MordFustang21/dbc"
```

### Basic Example

```go
func Divide(a, b int) int {
    // Precondition: divisor must not be zero
    dbc.Require(b != 0, "division by zero")
    
    result := a / b
    
    // Postcondition: division by 1 should return the original number
    dbc.Ensure(b != 1 || result == a, "identity property violated")
    
    return result
}
```

### Contract Types

#### Preconditions: `Require()`

Use at the beginning of functions to validate inputs or state before execution:

```go
func ProcessPositiveNumber(n int) {
    dbc.Require(n > 0, "n must be positive, got %d", n)
    // ...process the number...
}
```

#### Postconditions: `Ensure()`

Use before returning to validate outputs or state changes:

```go
func GetAbsoluteValue(n int) int {
    result := n
    if n < 0 {
        result = -n
    }
    
    dbc.Ensure(result >= 0, "absolute value must be non-negative, got %d", result)
    return result
}
```

#### Invariants: `Invariant()`

Use to validate object or system state that should always hold true:

```go
func (c *Counter) Increment(){
    dbc.Invariant(c.count >= 0, "counter should never be negative")
    c.count++
    dbc.Invariant(c.count > 0, "counter must be positive after increment")
}
```

#### Lazy Checks: `LazyRequire(), LazyEnsure(), LazyInvariant(), LazyCheck()`
Use for more complex conditions that may be expensive to evaluate. This is because the compiler will
not inline these functions, allowing them to be evaluated only when necessary.

```go
func ProcessData(data []int) {
	dbc.LazyRequire(func() bool {
		return len(data) > 0
	}, "data cannot be empty")
	
	// Process data...
	
	dbc.LazyEnsure(func() bool {
		return len(data) < 1000
	}, "data size should be less than 1000")
}
```


#### Always-On Checks: `Check()`

Use for critical assertions that should run even in production builds:

```go
func GetItem(slice []int, index int) int {
    dbc.Check(index >= 0 && index < len(slice), "index out of bounds: %d (len=%d)", index, len(slice))
    return slice[index]
}
```

## Production Builds

In production mode, `Require()`, `Ensure()`, and `Invariant()` become no-ops:

```bash
# Build with DBC checks enabled (development/testing)
go build

# Build with DBC checks disabled (production) - only Check() remains active
go build -tags production
```

## Why Use Design by Contract?

- **Self-documenting code**: Contracts express expectations directly in code
- **Early bug detection**: Failures occur close to the cause, not where symptoms appear 
- **Better testing**: Contract violations indicate test gaps
- **Robust interfaces**: Clarifies responsibility between caller and implementation
- **Zero production overhead**: Remove checks in production without changing code
