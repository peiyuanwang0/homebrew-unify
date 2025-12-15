//go:build linux && arm64
// +build linux,arm64

package main

import "embed"

//go:embed tools/linux_arm64/*
var FS embed.FS