//go:build !default_config

package main

import _ "embed"

//go:embed config_placeholder.bin
var configBytes []byte
