//go:build production

// Package dbc provides simple Design by Contract assertions (preconditions,
// postconditions, invariants) that are active during development and testing,
// but become no-ops when the "production" build tag is used.
// This file provides the no-op implementations for production builds.
package dbc

import "fmt" // Needed only for Check

// Require is a no-op in production builds.
func Require(condition bool, msgFormat string, args ...any) {
	// No-op
}

// Ensure is a no-op in production builds.
func Ensure(condition bool, msgFormat string, args ...any) {
	// No-op
}

// Invariant is a no-op in production builds.
func Invariant(condition bool, msgFormat string, args ...any) {
	// No-op
}

// Check executes the assertion regardless of build tags. It panics if the condition is false.
// It remains active even in production builds for essential internal checks.
// NOTE: This function *always* runs, unlike Require/Ensure/Invariant.
func Check(condition bool, msgFormat string, args ...any) {
	if !condition {
		// Simple panic message for production Check, avoiding runtime dependency.
		panic(fmt.Sprintf("Check violation: "+msgFormat, args...))
	}
}

// LazyRequire allows lazy evaluation of a precondition.
// It accepts a closure that returns a boolean and evaluates it only if needed.
func LazyRequire(condition func() bool, msgFormat string, args ...any) {
	// No-op in production builds
}

// LazyEnsure allows lazy evaluation of a postcondition.
func LazyEnsure(condition func() bool, msgFormat string, args ...any) {
	// No-op in production builds
}

// LazyInvariant allows lazy evaluation of an invariant.
func LazyInvariant(condition func() bool, msgFormat string, args ...any) {
	// No-op in production builds
}

// LazyCheck allows lazy evaluation of a condition for internal checks.
func LazyCheck(condition func() bool, msgFormat string, args ...any) {
	if !condition() {
		// Simple panic message for production LazyCheck, avoiding runtime dependency.
		panic(fmt.Sprintf("LazyCheck violation: "+msgFormat, args...))
	}
}
