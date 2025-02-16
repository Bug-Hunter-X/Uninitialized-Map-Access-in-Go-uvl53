# Uninitialized Map Access in Go

This repository demonstrates a common, yet subtle, error in Go related to accessing keys in uninitialized maps.  Unlike some languages that would throw an exception, Go silently returns the zero value for the map's value type when you attempt to access a key that doesn't exist in an uninitialized map. This behavior can lead to difficult-to-debug issues.

**The Problem:**

The `bug.go` file shows a simple example.  The map `m` is declared but not initialized. Accessing `m["a"]` returns 0 (the zero value for `int`), without raising a panic or error, making the bug hard to spot.

**The Solution:**

The `bugSolution.go` file demonstrates the correct way to handle this situation by checking if the key exists before accessing it or by explicitly initializing the map.