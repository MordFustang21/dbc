//go:build !production

// Package dbc provides simple Design by Contract assertions (preconditions,
// postconditions, invariants) that are active during development and testing,
// but become no-ops when the "production" build tag is used.
package dbc

import (
	"fmt"
	"runtime"
)

// panicWithCaller formats the message and panics, including caller info.
func panicWithCaller(contractType string, skip int, msgFormat string, args ...any) {
	// Get caller information (skip panicWithCaller and the dbc function itself)
	pc, file, line, ok := runtime.Caller(skip + 1) // +1 for this helper func
	callerInfo := "unknown:0"
	if ok {
		fn := runtime.FuncForPC(pc)
		callerFunc := "unknown()"
		if fn != nil {
			callerFunc = fn.Name() + "()"
		}
		callerInfo = fmt.Sprintf("%s:%d", file, line)
		msgFormat = fmt.Sprintf("%s violation in %s [%s]: %s", contractType, callerFunc, callerInfo, msgFormat)
	} else {
		msgFormat = fmt.Sprintf("%s violation [caller unknown]: %s", contractType, msgFormat)
	}

	panic(fmt.Sprintf(msgFormat, args...))
}

// Require checks for a precondition.
// It panics if the condition is false.
// Use this at the beginning of a function to validate inputs or state.
// No-op if built with "-tags production".
func Require(condition bool, msgFormat string, args ...any) {
	if !condition {
		panicWithCaller("Precondition", 2, msgFormat, args...) // skip=2 (Require + panicWithCaller)
	}
}

// Ensure checks for a postcondition.
// It panics if the condition is false.
// Use this typically with defer at the beginning of a function
// (or just before returning) to validate outputs or state changes.
// No-op if built with "-tags production".
func Ensure(condition bool, msgFormat string, args ...any) {
	if !condition {
		panicWithCaller("Postcondition", 2, msgFormat, args...) // skip=2 (Ensure + panicWithCaller)
	}
}

// Invariant checks for an invariant condition.
// It panics if the condition is false.
// Use this to check conditions that should always hold true for an object
// or module state, typically at the start and end of public methods.
// No-op if built with "-tags production".
func Invariant(condition bool, msgFormat string, args ...any) {
	if !condition {
		panicWithCaller("Invariant", 2, msgFormat, args...) // skip=2 (Invariant + panicWithCaller)
	}
}

// Check executes the assertion regardless of build tags, primarily for internal checks
// within the dbc package itself or situations where the check must always run.
// It panics if the condition is false.
// NOTE: This function *always* runs, unlike Require/Ensure/Invariant.
func Check(condition bool, msgFormat string, args ...any) {
	if !condition {
		panicWithCaller("Check", 2, msgFormat, args...) // skip=2 (Check + panicWithCaller)
	}
}

// LazyRequire allows lazy evaluation of a precondition.
// It accepts a closure that returns a boolean and evaluates it only if needed.
func LazyRequire(condition func() bool, msgFormat string, args ...any) {
	if !condition() {
		panicWithCaller("Precondition", 2, msgFormat, args...)
	}
}

// LazyEnsure allows lazy evaluation of a postcondition.
func LazyEnsure(condition func() bool, msgFormat string, args ...any) {
	if !condition() {
		panicWithCaller("Postcondition", 2, msgFormat, args...)
	}
}

// LazyInvariant allows lazy evaluation of an invariant.
func LazyInvariant(condition func() bool, msgFormat string, args ...any) {
	if !condition() {
		panicWithCaller("Invariant", 2, msgFormat, args...)
	}
}

// LazyCheck allows lazy evaluation of a condition for internal checks.
func LazyCheck(condition func() bool, msgFormat string, args ...any) {
	if !condition() {
		panicWithCaller("Check", 2, msgFormat, args...)
	}
}
