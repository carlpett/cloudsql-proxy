// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logging contains helpers to support log messages. If you are using
// the Cloud SQL Auth proxy as a Go library, you can override these variables to
// control where log messages end up.
package logging

import (
	"log"
	"os"

	"go.uber.org/zap"
)

// Verbosef is called to write verbose logs, such as when a new connection is
// established correctly.
var Verbosef = log.Printf

// Infof is called to write informational logs, such as when startup has
var Infof = log.Printf

// Errorf is called to write an error log, such as when a new connection fails.
var Errorf = log.Printf

// LogDebugToStdout updates Verbosef and Info logging to use stdout instead of stderr.
func LogDebugToStdout() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	Verbosef = logger.Printf
	Infof = logger.Printf
}

func noop(string, ...interface{}) {}

// LogVerboseToNowhere updates Verbosef so verbose log messages are discarded
func LogVerboseToNowhere() {
	Verbosef = noop
}

// DisableLogging sets all logging levels to no-op's.
func DisableLogging() {
	Verbosef = noop
	Infof = noop
	Errorf = noop
}

// EnableStructuredLogs replaces all logging functions with structured logging
// variants.
func EnableStructuredLogs() (func(), error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return func() {}, err
	}
	sugar := logger.Sugar()
	Verbosef = sugar.Debugf
	Infof = sugar.Infof
	Errorf = sugar.Errorf
	return func() {
		logger.Sync()
	}, nil
}
