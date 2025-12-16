package natsgo

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Handler func(*nats.Msg)
type NatsClient struct {
	conn *nats.Conn
}

func New(name, url string, timeout time.Duration, retries int, pingInterval time.Duration) (*NatsClient, error) {
	var nc *nats.Conn
	var err error
	for i := range retries {
		nc, err = nats.Connect(url,
			nats.Name(name),
			nats.Timeout(timeout),
			nats.PingInterval(pingInterval),
			nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
				log.Printf("[NATS] Disconnected from %s", c.ConnectedUrl())
			}),
		)
		if err == nil {
			break
		}

		if i < retries {
			log.Printf("[NATS] Failed to connect (attempt %d/%d): %v", i+1, retries+1, err)
			time.Sleep(2 * time.Second)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("nats: failed to connect: %w", err)
	}
	return &NatsClient{
		conn: nc,
	}, nil
}

func (nc *NatsClient) Publish(subj string, data []byte) error {
	err := nc.conn.Publish(subj, data)
	if err != nil {
		return fmt.Errorf("error when publishing %s : %w", subj, err)
	}
	return nil
}

func (nc *NatsClient) Subscribe(subj string, handler nats.MsgHandler) (*nats.Subscription, error) {
	sub, err := nc.conn.Subscribe(subj, handler)
	if err != nil {
		return nil, fmt.Errorf("error subscribing to subject %s: %w", subj, err)
	}
	return sub, nil
}

func (nc *NatsClient) Close() {
	nc.conn.Close()
}
