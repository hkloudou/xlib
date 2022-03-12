package xtransport

import (
	"crypto/tls"
	"io"
	"net"
	"time"
)

type tcpSocket[T Writer] struct {
	net.Conn
	timeout time.Duration
	*session
	obound chan T
	closed bool
}

func (t *tcpSocket[T]) ConnectState() *tls.ConnectionState {
	if c2, ok := t.Conn.(*tls.Conn); ok {
		tmp := c2.ConnectionState()
		return &tmp
	}
	return nil
}

func (t *tcpSocket[T]) Local() string {
	return t.Conn.LocalAddr().String()
}

func (t *tcpSocket[T]) Remote() string {
	return t.Conn.RemoteAddr().String()
}

func (t *tcpSocket[T]) Recv(fc func(r io.Reader) (T, error)) (T, error) {
	if t.timeout > time.Duration(0) {
		t.Conn.SetDeadline(time.Now().Add(t.timeout))
	}
	return fc(t.Conn)
}

func (t *tcpSocket[T]) loop() {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	for v := range t.obound {
		if t.timeout > time.Duration(0) {
			t.Conn.SetDeadline(time.Now().Add(t.timeout))
		}
		if err := v.Write(t.Conn); err != nil {
			t.Close()
			return
		}
	}
}

func (t *tcpSocket[T]) Send(m T) error {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	select {
	case t.obound <- m:
	case <-time.After(5 * time.Second):
		t.Close()
		break
	}
	return nil
	// return m.Write(t.Conn)
}

func (t *tcpSocket[T]) SetTimeOut(duration time.Duration) {
	t.timeout = duration
}

func (t *tcpSocket[T]) Close() error {
	if t.closed {
		return nil
	}
	t.closed = true
	close(t.obound)
	return t.Conn.Close()
}
