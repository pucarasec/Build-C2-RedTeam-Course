//go:build default_config

package main

import _ "embed"

//go:embed config.json
var configBytes []byte
