package gins

import (
	"context"
	"testing"
)

func TestConfigureRedisConfigure(t *testing.T) {
	conf := ConfigureRedisConnection(map[string]interface{}{
		"network": "unix",
		"addr":    "addr",
	})

	if conf.Network != "unix" {
		t.Errorf("unexpected network, expecting unix but %s", conf.Network)
	}

	if conf.Addr != "addr" {
		t.Errorf("unexpected address, expecting addr but %s", conf.Addr)
	}
}

func TestConnectRedis(t *testing.T) {
	conf := ConfigureRedisConnection(nil)
	cnx := ConnectRedis(conf)

	if cnx == nil {
		t.Error("no connection")
	} else if pong := cnx.Ping(context.Background()); pong.Err() != nil {
		t.Error(pong.Err())
	}

}
