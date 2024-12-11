# Go Append to Uninitialized Slice
This repository demonstrates an uncommon error in Go: attempting to append to an uninitialized slice.  In Go, slices are dynamically sized, but they must be initialized before use. This example showcases the resulting runtime panic and offers a corrected version.

## Bug
The bug lies in the `main` function.  The variable `x` is declared as an integer, but the code attempts to append to it as if it were a slice. This leads to a runtime panic.

## Solution
The solution initializes a slice before appending elements, correctly demonstrating the proper way to use slices in Go.