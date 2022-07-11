// Copyright 2021 Gravitational, Inc
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

package teleterm

import (
	"os"

	"github.com/gravitational/teleport/lib/utils"

	"github.com/gravitational/trace"
)

// Config describes teleterm configuration
type Config struct {
	// Addr is the bind address for the server
	Addr string
	// ShutdownSignals is the set of captured signals that cause server shutdown.
	ShutdownSignals []os.Signal
	// HomeDir is the directory to store cluster profiles
	HomeDir string
	// CertsDir is the directory where the certs are stored
	CertsDir string
	// InsecureSkipVerify is an option to skip HTTPS cert check
	InsecureSkipVerify bool
}

// CheckAndSetDefaults checks and sets default config values.
func (c *Config) CheckAndSetDefaults() error {
	if c.HomeDir == "" {
		return trace.BadParameter("missing home directory")
	}

	if c.CertsDir == "" {
		return trace.BadParameter("missing certs directory")
	}

	if c.Addr == "" {
		return trace.BadParameter("missing network address")
	}

	addr, err := utils.ParseAddr(c.Addr)
	if err != nil {
		return trace.Wrap(err)
	}

	if !(addr.Network() == "unix" || addr.Network() == "tcp") {
		return trace.BadParameter("only unix and tcp sockets are supported")
	}

	return nil
}
