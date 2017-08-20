package version

import (
	"golang.org/x/net/context"

	log "github.com/sirupsen/logrus"
)

var (
	// Version : Git version identifier generated via `git describe --always`
	Version = "dev"

	// Build : Timestamp when the application was build
	Build = "N/A"
)

// Server : Implementation of the gRPC service
type Server struct{}

/************************************************************************************
 * Start gRPC Service implementation
 */

// Version : Returns the Server version
func (s Server) Version(ctx context.Context, in *VersionMessage) (*VersionMessage, error) {
	log.WithField("client_version", in.Version).
		WithField("client_build", in.Build).
		WithField("call", "Version").Print("Incoming gRPC call")
	return &VersionMessage{Version: Version, Build: Build}, nil
}
