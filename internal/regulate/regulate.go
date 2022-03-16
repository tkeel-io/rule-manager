package regulate

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

/*

	fps调节器，降低刷新频率

*/

var defaultExpiration, cleanupInterval time.Duration = 2 * time.Second, 1 * time.Second

type Regulation struct {
	cache *gocache.Cache
}

func NewRegulation() *Regulation {
	return &Regulation{
		cache: gocache.New(defaultExpiration, cleanupInterval),
	}
}

func (this *Regulation) Regulate(key string, handle func()) bool {
	_, has := this.cache.Get(key)
	if !has {
		this.cache.SetDefault(key, nil)
		handle()
	}
	return has
}
