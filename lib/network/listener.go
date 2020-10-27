package network

import (
	"context"
	"net"

	"github.com/pkg/errors"
)

// ListenerOptions
type ListenerOptions struct {
	HostPort string
}

type listener struct {
	net.Listener
}

// NewTCPListener create tcp listener to provided port
func NewTCPListener(ctx context.Context, opts ListenerOptions) (*listener, error) {

	lis, err := net.Listen("tcp", opts.HostPort)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen tcp on:%v", opts.HostPort)
	}

	return &listener{Listener: lis}, nil
}
