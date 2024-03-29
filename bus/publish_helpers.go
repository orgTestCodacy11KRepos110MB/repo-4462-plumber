package bus

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
)

// PublishCreateConnection publishes a CreateConnection message, which other plumber instances will receive
// and add the connection to their local in-memory maps
func (b *Bus) PublishCreateConnection(ctx context.Context, conn *opts.ConnectionOptions) error {
	return b.publishConnectionMessage(ctx, CreateConnection, conn)
}

// PublishUpdateConnection publishes an UpdateConnection message, which other plumber instances will receive
// and update the connection in their local in-memory maps
func (b *Bus) PublishUpdateConnection(ctx context.Context, conn *opts.ConnectionOptions) error {
	return b.publishConnectionMessage(ctx, UpdateConnection, conn)
}

// PublishDeleteConnection publishes a DeleteConnection message, which other plumber instances will receive
// and delete from their local in-memory maps
func (b *Bus) PublishDeleteConnection(ctx context.Context, conn *opts.ConnectionOptions) error {
	return b.publishConnectionMessage(ctx, DeleteConnection, conn)
}

// PublishCreateRelay publishes a CreateRelay message, which other plumber instances will receive
// and add the service to their local in-memory maps
func (b *Bus) PublishCreateRelay(ctx context.Context, relay *opts.RelayOptions) error {
	return b.publishRelayMessage(ctx, CreateRelay, relay)
}

// PublishUpdateRelay publishes an UpdateRelay message, which other plumber instances will receive
// and update the connection in their local in-memory maps
func (b *Bus) PublishUpdateRelay(ctx context.Context, relay *opts.RelayOptions) error {
	return b.publishRelayMessage(ctx, UpdateRelay, relay)
}

// PublishDeleteRelay publishes a DeleteRelay message, which other plumber instances will receive
// and delete from their local in-memory maps
func (b *Bus) PublishDeleteRelay(ctx context.Context, relay *opts.RelayOptions) error {
	return b.publishRelayMessage(ctx, DeleteRelay, relay)
}

// PublishStopRelay broadcasts a StopRelay message which will cause all plumber
// instances to stop the relay and remove it from their in-memory cache.
func (b *Bus) PublishStopRelay(ctx context.Context, relay *opts.RelayOptions) error {
	return b.publishRelayMessage(ctx, StopRelay, relay)
}

// PublishResumeRelay broadcasts a ResumeRelay message which will cause all plumber
// instances to start a stopped relay and add it to their in-memory cache.
func (b *Bus) PublishResumeRelay(ctx context.Context, relay *opts.RelayOptions) error {
	return b.publishRelayMessage(ctx, ResumeRelay, relay)
}

// PublishCreateTunnel publishes a CreateTunnel message, which other plumber instances will receive
// and add the service to their local in-memory maps
func (b *Bus) PublishCreateTunnel(ctx context.Context, tunnelOptions *opts.TunnelOptions) error {
	return b.publishTunnelMessage(ctx, CreateTunnel, tunnelOptions)
}

// PublishUpdateTunnel publishes an UpdateTunnel message, which other plumber instances will receive
// and update the connection in their local in-memory maps
func (b *Bus) PublishUpdateTunnel(ctx context.Context, tunnelOptions *opts.TunnelOptions) error {
	return b.publishTunnelMessage(ctx, UpdateTunnel, tunnelOptions)
}

// PublishDeleteTunnel publishes a DeleteTunnel message, which other plumber instances will receive
// and delete from their local in-memory maps
func (b *Bus) PublishDeleteTunnel(ctx context.Context, tunnelOptions *opts.TunnelOptions) error {
	return b.publishTunnelMessage(ctx, DeleteTunnel, tunnelOptions)
}

// PublishStopTunnel broadcasts a StopTunnel message which will cause all plumber
// instances to stop the relay and remove it from their in-memory cache.
func (b *Bus) PublishStopTunnel(ctx context.Context, tunnelOptions *opts.TunnelOptions) error {
	return b.publishTunnelMessage(ctx, StopTunnel, tunnelOptions)
}

// PublishResumeTunnel broadcasts a ResumeTunnel message which will cause all plumber
// instances to start a stopped relay and add it to their in-memory cache.
func (b *Bus) PublishResumeTunnel(ctx context.Context, tunnelOptions *opts.TunnelOptions) error {
	return b.publishTunnelMessage(ctx, ResumeTunnel, tunnelOptions)
}

func (b *Bus) publishConnectionMessage(ctx context.Context, action Action, conn *opts.ConnectionOptions) error {
	data, err := proto.Marshal(conn)
	if err != nil {
		return errors.Wrapf(err, "unable to marshal connection message for '%s'", conn.XId)
	}

	return b.broadcast(ctx, &Message{
		Action:    action,
		Data:      data,
		EmittedBy: b.config.PersistentConfig.PlumberID,
		EmittedAt: time.Now().UTC(),
	})
}

func (b *Bus) publishRelayMessage(ctx context.Context, action Action, relay *opts.RelayOptions) error {
	data, err := proto.Marshal(relay)
	if err != nil {
		return errors.Wrapf(err, "unable to marshal relay message for '%s'", relay.XRelayId)
	}

	return b.broadcast(ctx, &Message{
		Action:    action,
		Data:      data,
		EmittedBy: b.config.PersistentConfig.PlumberID,
		EmittedAt: time.Now().UTC(),
	})
}

func (b *Bus) publishTunnelMessage(ctx context.Context, action Action, tunnelOptions *opts.TunnelOptions) error {
	data, err := proto.Marshal(tunnelOptions)
	if err != nil {
		return errors.Wrapf(err, "unable to marshal tunnel message for '%s'", tunnelOptions.XTunnelId)
	}

	return b.broadcast(ctx, &Message{
		Action:    action,
		Data:      data,
		EmittedBy: b.config.PersistentConfig.PlumberID,
		EmittedAt: time.Now().UTC(),
	})
}
