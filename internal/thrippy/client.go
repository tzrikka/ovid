// Package thrippy provides minimal and lightweight wrappers for some client
// functionalities of the [Thrippy gRPC service]. It is meant to facilitate
// code reuse, not to provide a complete native layer on top of gRPC.
//
// [Thrippy gRPC service]: https://github.com/tzrikka/thrippy-api/blob/main/proto/thrippy/v1/thrippy.proto
package thrippy

import (
	"context"
	"time"

	"github.com/lithammer/shortuuid/v4"
	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"

	thrippypb "github.com/tzrikka/thrippy-api/thrippy/v1"
)

const (
	timeout = 3 * time.Second
)

type LinkClient struct {
	LinkID   string
	grpcAddr string
	creds    credentials.TransportCredentials
}

func NewLinkClient(linkID string, cmd *cli.Command) LinkClient {
	return LinkClient{
		LinkID:   linkID,
		grpcAddr: cmd.String("thrippy-server-addr"),
		creds:    secureCreds(cmd),
	}
}

func (t *LinkClient) connection(l log.Logger, providerName string) (*grpc.ClientConn, error) {
	if t.LinkID == "" {
		msg := "Thrippy link ID not configured for " + providerName
		l.Warn(msg)
		return nil, temporal.NewNonRetryableApplicationError(msg, "error", nil)
	}

	if _, err := shortuuid.DefaultEncoder.Decode(t.LinkID); err != nil {
		msg := "invalid Thrippy link ID configured for " + providerName
		l.Warn(msg, "link_id", t.LinkID)
		return nil, temporal.NewNonRetryableApplicationError(msg, "error", nil, t.LinkID)
	}

	conn, err := grpc.NewClient(t.grpcAddr, grpc.WithTransportCredentials(t.creds))
	if err != nil {
		l.Error("failed to create gRPC client connection", "error", err.Error(), "grpc_addr", t.grpcAddr)
	}

	return conn, err
}

func (t *LinkClient) LinkCreds(ctx context.Context, providerName string) (map[string]string, error) {
	l := activity.GetLogger(ctx)

	conn, err := t.connection(l, providerName)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := thrippypb.NewThrippyServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resp, err := c.GetCredentials(ctx, thrippypb.GetCredentialsRequest_builder{
		LinkId: proto.String(t.LinkID),
	}.Build())
	if err != nil {
		l.Error("Thrippy GetCredentials error", "error", err.Error(), "link_id", t.LinkID)
		return nil, err
	}

	return resp.GetCredentials(), nil
}

func (t *LinkClient) LinkData(ctx context.Context, providerName string) (string, map[string]string, error) {
	l := activity.GetLogger(ctx)

	conn, err := t.connection(l, providerName)
	if err != nil {
		return "", nil, err
	}
	defer conn.Close()

	c := thrippypb.NewThrippyServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Template.
	resp1, err := c.GetLink(ctx, thrippypb.GetLinkRequest_builder{
		LinkId: proto.String(t.LinkID),
	}.Build())
	if err != nil {
		l.Error("Thrippy GetLink error", "error", err.Error(), "link_id", t.LinkID)
		return "", nil, err
	}

	// Credentials.
	resp2, err := c.GetCredentials(ctx, thrippypb.GetCredentialsRequest_builder{
		LinkId: proto.String(t.LinkID),
	}.Build())
	if err != nil {
		l.Error("Thrippy GetCredentials error", "error", err.Error(), "link_id", t.LinkID)
		return "", nil, err
	}

	return resp1.GetTemplate(), resp2.GetCredentials(), nil
}
