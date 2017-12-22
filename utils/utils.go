package utils

import (
	"errors"

	"github.com/xurwxj/cryptkv/backend"
	"github.com/xurwxj/cryptkv/backend/consul"
	"github.com/xurwxj/cryptkv/backend/etcd"
	"github.com/xurwxj/cryptkv/backend/zookeeper"
)

func GetPlainConf(backend, endpoint, key string) ([]byte, error) {
	backendStore, err := GetBackendStore(backend, endpoint)
	if err != nil {
		return nil, err
	}
	v, err := GetPlain(key, backendStore)
	if err != nil {
		return nil, err
	}
	return v, err
}

func GetPlain(key string, store backend.Store) ([]byte, error) {
	var value []byte
	data, err := store.Get(key)
	if err != nil {
		return value, err
	}
	return data, err
}

func GetBackendStore(provider, endpoint string) (backend.Store, error) {
	if endpoint == "" {
		switch provider {
		case "consul":
			endpoint = "127.0.0.1:8500"
		case "etcd":
			endpoint = "http://127.0.0.1:4001"
		case "zookeeper":
			endpoint = "127.0.0.1:2181"
		}
	}
	machines := []string{endpoint}
	switch provider {
	case "etcd":
		return etcd.New(machines)
	case "consul":
		return consul.New(machines)
	case "zookeeper":
		return zookeeper.New(machines)
	default:
		return nil, errors.New("invalid backend " + provider)
	}
}
