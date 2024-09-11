# EZKL FFI

C and CGO bindings for the EZKL library.

## Requirements

1. **Go and `c-for-go`:**
   - Install Go.
   - Install `c-for-go` using the following command:
     ```bash
     go install github.com/xlab/c-for-go@latest
     ```
   - Ensure that `$GOPATH/bin` is added to your `$PATH`.

2. **Rust Toolchain:**
   - With the actual version, we require the nightly Rust toolchain to be installed.

## Setup

1. To compile the library and copy it to `LIBRARY_PATH` so the library can be found by compilers, run:
   ```bash
   make

2. To clean up the generated files, run:
   ```bash
   make clean
