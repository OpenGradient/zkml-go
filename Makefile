# Determine the OS
UNAME_S := $(shell uname -s)

# Rust toolchain and target
ifeq ($(UNAME_S),Darwin)
    RUST_TOOLCHAIN := nightly-2024-04-17
    RUST_TARGET := aarch64-apple-darwin
else ifeq ($(UNAME_S),Linux)
    RUST_TOOLCHAIN := nightly
    RUST_TARGET := x86_64-unknown-linux-gnu
else
    $(error Unsupported OS: $(UNAME_S))
endif

# variables
RUST_TARGET_DIR := ./rust/target/$(RUST_TARGET)/release
CGO_DIR := ./cgo
LIB_INSTALL_DIR := /usr/local/lib
DEPS:=ezkl-ffi.h libezkl-ffi.a

all: $(DEPS)
.PHONY: all

$(DEPS): .install-ezkl-ffi  ;

# Compile the Rust library and generate the required files
.install-ezkl-ffi: rust
	rustup override set $(RUST_TOOLCHAIN)
	cd rust && cargo build --release --target $(RUST_TARGET) --all; cd ..
	find $(RUST_TARGET_DIR) -type f -name "ezkl-ffi.h" -print0 | xargs -0 ls -t | head -n 1 | xargs -I {} cp {} $(CGO_DIR)/ezkl-ffi.h
	find $(RUST_TARGET_DIR) -type f -name "libezkl_ffi.a" -print0 | xargs -0 ls -t | head -n 1 | xargs -I {} cp {} $(CGO_DIR)/libezkl_ffi.a
	find $(RUST_TARGET_DIR) -type f -name "libezkl_ffi.a" -print0 | xargs -0 ls -t | head -n 1 | xargs -I {} sudo cp {} $(LIB_INSTALL_DIR)/libezkl_ffi.a
	c-for-go --ccincl ezkl-ffi.yml
	@touch $@

clean:
	rm -rf $(DEPS) .install-ezkl-ffi
	rm -rf $(RUST_TARGET_DIR)/build/ezkl-ffi-*
	rm -rf $(CGO_DIR)/*.go
	rm -rf $(CGO_DIR)/*.h
	rm -rf $(CGO_DIR)/*.a
	sudo rm -f $(LIB_INSTALL_DIR)/libezkl_ffi.a
.PHONY: clean