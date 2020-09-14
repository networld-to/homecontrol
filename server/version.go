package main

import (
	"context"

	"github.com/networld-to/homecontrol/api/generated/version"
	"github.com/networld-to/homecontrol/utils"
	log "github.com/sirupsen/logrus"
)

// VersionServer : Implementation of the gRPC service
type VersionServer struct{}

/************************************************************************************
 * Start gRPC Service implementation
 */

// Version : Returns the Server version
func (s VersionServer) Version(ctx context.Context, in *version.VersionMessage) (*version.VersionMessage, error) {
	log.WithField("client_version", in.Version).
		WithField("client_build", in.Build).
		WithField("call", "Version").Print("Incoming gRPC call")
	return &version.VersionMessage{Version: utils.Version, Build: utils.Build}, nil
}
