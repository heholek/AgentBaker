// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package main

import (
	"os"

	"github.com/Azure/agentbaker/cmd"
	colorable "github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	if err := cmd.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
