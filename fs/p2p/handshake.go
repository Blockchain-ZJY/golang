package p2p

import "errors"

var ErrInvalidHandshake = errors.New("invalid handshak")

type HandShakeFunc func(any) error

func NOPHandshakeFunc(any) error {
	return nil
}
