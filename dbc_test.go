package dbc

import (
	"fmt"
	"strings"
	"testing"
)

// TestRequire tests the Require function
func TestRequire(t *testing.T) {
	// Test that valid conditions don't panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Require panicked unexpectedly: %v", r)
			}
		}()
		Require(true, "This should not panic")
	}()

	// Test that invalid conditions panic with correct message
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Require failed to panic when condition was false")
			} else {
				msg := fmt.Sprint(r)
				if !strings.Contains(msg, "Precondition violation") {
					t.Errorf("Wrong panic message: %v", msg)
				}
				if !strings.Contains(msg, "TestRequire") {
					t.Errorf("Panic message didn't contain caller info: %v", msg)
				}
			}
		}()
		Require(false, "Test message %d", 42)
	}()
}

// TestEnsure tests the Ensure function
func TestEnsure(t *testing.T) {
	// Test that valid conditions don't panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Ensure panicked unexpectedly: %v", r)
			}
		}()
		Ensure(true, "This should not panic")
	}()

	// Test that invalid conditions panic with correct message
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Ensure failed to panic when condition was false")
			} else {
				msg := fmt.Sprint(r)
				if !strings.Contains(msg, "Postcondition violation") {
					t.Errorf("Wrong panic message: %v", msg)
				}
				if !strings.Contains(msg, "TestEnsure") {
					t.Errorf("Panic message didn't contain caller info: %v", msg)
				}
			}
		}()
		Ensure(false, "Test message %d", 42)
	}()
}

// TestInvariant tests the Invariant function
func TestInvariant(t *testing.T) {
	// Test that valid conditions don't panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Invariant panicked unexpectedly: %v", r)
			}
		}()
		Invariant(true, "This should not panic")
	}()

	// Test that invalid conditions panic with correct message
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Invariant failed to panic when condition was false")
			} else {
				msg := fmt.Sprint(r)
				if !strings.Contains(msg, "Invariant violation") {
					t.Errorf("Wrong panic message: %v", msg)
				}
				if !strings.Contains(msg, "TestInvariant") {
					t.Errorf("Panic message didn't contain caller info: %v", msg)
				}
			}
		}()
		Invariant(false, "Test message %d", 42)
	}()
}

// TestCheck tests the Check function
func TestCheck(t *testing.T) {
	// Test that valid conditions don't panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Check panicked unexpectedly: %v", r)
			}
		}()
		Check(true, "This should not panic")
	}()

	// Test that invalid conditions panic with correct message
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Check failed to panic when condition was false")
			} else {
				msg := fmt.Sprint(r)
				if !strings.Contains(msg, "Check violation") {
					t.Errorf("Wrong panic message: %v", msg)
				}
				if !strings.Contains(msg, "TestCheck") {
					t.Errorf("Panic message didn't contain caller info: %v", msg)
				}
			}
		}()
		Check(false, "Test message %d", 42)
	}()
}

// Here are examples in godoc format

// Example shows basic usage of the dbc package.
func Example() {
	// Create a simple function with design-by-contract principles
	divide := func(a, b int) int {
		// Precondition: divisor must not be zero
		Require(b != 0, "divisor must not be zero")
		
		result := a / b
		
		// Postcondition: if b is 1, result must equal a
		Ensure(b != 1 || result == a, "identity property not preserved")
		
		return result
	}

	// This will work fine
	result := divide(10, 2)
	fmt.Println("10 / 2 =", result)
	
	// Uncomment to see precondition failure:
	// divide(5, 0)
	
	// Output: 10 / 2 = 5
}

// ExampleRequire demonstrates how to use preconditions
func ExampleRequire() {
	processPositiveNumber := func(n int) {
		Require(n > 0, "n must be positive, got %d", n)
		// Process the number...
		fmt.Println("Processing:", n)
	}
	
	processPositiveNumber(5)
	
	// Uncommenting this would cause a panic:
	// processPositiveNumber(-1)
	
	// Output: Processing: 5
}

// ExampleEnsure demonstrates how to use postconditions
func ExampleEnsure() {
	getAbsoluteValue := func(n int) int {
		result := n
		if n < 0 {
			result = -n
		}
		
		// Verify our result is non-negative
		Ensure(result >= 0, "absolute value must be non-negative, got %d", result)
		return result
	}
	
	fmt.Println(getAbsoluteValue(-10))
	fmt.Println(getAbsoluteValue(5))
	
	// Output:
	// 10
	// 5
}

// ExampleInvariant demonstrates how to use invariants
func ExampleInvariant() {
	// Simple counter that should never go negative
	counter := 0
	
	increment := func() {
		Invariant(counter >= 0, "counter should never be negative")
		counter++
		Invariant(counter >= 0, "counter should never be negative")
		fmt.Println("Counter:", counter)
	}
	
	decrement := func() {
		Invariant(counter >= 0, "counter should never be negative")
		if counter > 0 {
			counter--
		}
		Invariant(counter >= 0, "counter should never be negative")
		fmt.Println("Counter:", counter)
	}
	
	increment() // 1
	increment() // 2
	decrement() // 1
	decrement() // 0
	
	// Output:
	// Counter: 1
	// Counter: 2
	// Counter: 1
	// Counter: 0
}

// ExampleCheck demonstrates how to use the always-on check function
func ExampleCheck() {
	validateSliceAccess := func(slice []int, index int) int {
		// This check will always run, even in production builds
		Check(index >= 0 && index < len(slice), "index out of bounds: %d (len=%d)", index, len(slice))
		return slice[index]
	}
	
	numbers := []int{10, 20, 30, 40, 50}
	value := validateSliceAccess(numbers, 2)
	fmt.Println("Value:", value)
	
	// This would panic:
	// validateSliceAccess(numbers, 10)
	
	// Output: Value: 30
}
