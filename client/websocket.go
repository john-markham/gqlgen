package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"strings"

	"github.com/john-markham/websocket"
)

const (
	connectionInitMsg = "connection_init" // Client -> Server
	startMsg          = "start"           // Client -> Server
	connectionAckMsg  = "connection_ack"  // Server -> Client
	connectionKaMsg   = "ka"              // Server -> Client
	dataMsg           = "data"            // Server -> Client
	errorMsg          = "error"           // Server -> Client
)

type operationMessage struct {
	Payload json.RawMessage `json:"payload,omitempty"`
	ID      string          `json:"id,omitempty"`
	Type    string          `json:"type"`
}

type Subscription struct {
	Close func() error
	Next  func(response any) error
}

func errorSubscription(err error) *Subscription {
	return &Subscription{
		Close: func() error { return nil },
		Next: func(response any) error {
			return err
		},
	}
}

func (p *Client) Websocket(query string, options ...Option) *Subscription {
	return p.WebsocketWithPayload(query, nil, options...)
}

// Grab a single response from a websocket based query
func (p *Client) WebsocketOnce(query string, resp any, options ...Option) error {
	sock := p.Websocket(query, options...)
	defer func() { _ = sock.Close() }()
	if reflect.ValueOf(resp).Kind() == reflect.Ptr {
		return sock.Next(resp)
	}
	// TODO: verify this is never called and remove it
	return sock.Next(&resp)
}

func (p *Client) WebsocketWithPayload(query string, initPayload map[string]any, options ...Option) *Subscription {
	r, err := p.newRequest(query, options...)
	if err != nil {
		return errorSubscription(fmt.Errorf("request: %w", err))
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return errorSubscription(fmt.Errorf("parse body: %w", err))
	}

	srv := httptest.NewServer(p.h)
	host := strings.ReplaceAll(srv.URL, "http://", "ws://")
	c, resp, err := websocket.DefaultDialer.Dial(host+r.URL.Path, r.Header)
	if err != nil {
		return errorSubscription(fmt.Errorf("dial: %w", err))
	}
	defer resp.Body.Close()

	initMessage := operationMessage{Type: connectionInitMsg}
	if initPayload != nil {
		initMessage.Payload, err = json.Marshal(initPayload)
		if err != nil {
			return errorSubscription(fmt.Errorf("parse payload: %w", err))
		}
	}

	if err = c.WriteJSON(initMessage); err != nil {
		return errorSubscription(fmt.Errorf("init: %w", err))
	}

	var ack operationMessage
	if err = c.ReadJSON(&ack); err != nil {
		return errorSubscription(fmt.Errorf("ack: %w", err))
	}

	if ack.Type != connectionAckMsg {
		return errorSubscription(fmt.Errorf("expected ack message, got %#v", ack))
	}

	var ka operationMessage
	if err = c.ReadJSON(&ka); err != nil {
		return errorSubscription(fmt.Errorf("ack: %w", err))
	}

	if ka.Type != connectionKaMsg {
		return errorSubscription(fmt.Errorf("expected ack message, got %#v", ack))
	}

	if err = c.WriteJSON(operationMessage{Type: startMsg, ID: "1", Payload: requestBody}); err != nil {
		return errorSubscription(fmt.Errorf("start: %w", err))
	}

	return &Subscription{
		Close: func() error {
			srv.Close()
			return c.Close()
		},
		Next: func(response any) error {
			for {
				var op operationMessage
				err := c.ReadJSON(&op)
				if err != nil {
					return err
				}

				switch op.Type {
				case dataMsg:
					break
				case connectionKaMsg:
					continue
				case errorMsg:
					return errors.New(string(op.Payload))
				default:
					return fmt.Errorf("expected data message, got %#v", op)
				}

				var respDataRaw Response
				err = json.Unmarshal(op.Payload, &respDataRaw)
				if err != nil {
					return fmt.Errorf("decode: %w", err)
				}

				// we want to unpack even if there is an error, so we can see partial responses
				unpackErr := unpack(respDataRaw.Data, response, p.dc)

				if respDataRaw.Errors != nil {
					return RawJsonError{respDataRaw.Errors}
				}
				return unpackErr
			}
		},
	}
}
