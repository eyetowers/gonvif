package gonvif

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/sync/singleflight"
)

type ClientPool interface {
	GetClient(ctx context.Context, baseURL, username, password string, verbose bool) (Client, error)
}

func NewPool(ttl time.Duration) ClientPool {
	cache := ttlcache.New(
		ttlcache.WithTTL[key, Client](ttl),
	)
	go cache.Start()

	return &pool{
		cache: cache,
	}
}

type key struct {
	baseURL  string
	username string
	password string
	verbose  bool
}

type pool struct {
	cache *ttlcache.Cache[key, Client]
	group singleflight.Group
}

func (p *pool) GetClient(
	ctx context.Context,
	baseURL, username, password string,
	verbose bool,
) (Client, error) {
	k := key{
		baseURL:  baseURL,
		username: username,
		password: password,
		verbose:  verbose,
	}
	item := p.cache.Get(k)

	if item != nil {
		return item.Value(), nil
	}

	return p.newClientSynced(ctx, k)
}

func (p *pool) newClientSynced(ctx context.Context, k key) (Client, error) {
	v, err, _ := p.group.Do(k.String(), func() (any, error) {
		return p.newClient(ctx, k)
	})
	if err != nil {
		return nil, err
	}
	return v.(Client), nil
}

func (p *pool) newClient(ctx context.Context, k key) (Client, error) {
	client, err := New(ctx, k.baseURL, k.username, k.password, k.verbose)
	if err != nil {
		return nil, err
	}
	p.cache.Set(k, client, ttlcache.DefaultTTL)
	return client, nil
}

var escaper = strings.NewReplacer(
	"\\", "\\\\",
	"|", "\\|",
)

func (k key) String() string {
	return fmt.Sprintf("%s|%s|%s|%v",
		escaper.Replace(k.baseURL),
		escaper.Replace(k.password),
		escaper.Replace(k.username),
		k.verbose)
}
