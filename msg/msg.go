// Copyright 2014 Akamai Technologies, Inc.
// Please see the accompanying LICENSE file for licensing information.

// Package msg defines the main messages used in the Focus protocol.
package msg

import (
	"github.com/mstone/focus/ot"
)

type OTServerMsg struct {
	Site int
	Rev  int
	Ack  bool
	Ops  ot.Ops
}

type OTClientMsg struct {
	Site int
	Rev  int
	Ops  ot.Ops
}
