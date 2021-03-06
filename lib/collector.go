/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package lib

import (
	"context"

	"github.com/loadimpact/k6/stats"
)

// Run Status used by cloud collector
const (
	RunStatusFinished           = 3
	RunStatusAbortedUser        = 5
	RunStatusAbortedSystem      = 6
	RunStatusAbortedScriptError = 7
	RunStatusAbortedThreshold   = 8
)

// A Collector abstracts the process of funneling samples to an external storage backend,
// such as an InfluxDB instance.
type Collector interface {
	// Init is called between the collector's creation and the call to Run().
	// You should do any lengthy setup here rather than in New.
	Init() error

	// Run is called in a goroutine and starts the collector. Should commit samples to the backend
	// at regular intervals and when the context is terminated.
	Run(ctx context.Context)

	// Collect receives a set of samples. This method is never called concurrently, and only while
	// the context for Run() is valid, but should defer as much work as possible to Run().
	Collect(samples []stats.SampleContainer)

	// Optionally return a link that is shown to the user.
	Link() string

	// Return the required system sample tags for the specific collector
	GetRequiredSystemTags() TagSet

	// Set run status
	SetRunStatus(status int)
}
