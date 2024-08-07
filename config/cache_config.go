package config

import (
	"sync"

	cacheconf "github.com/IacopoMelani/Go-Starter-Project/pkg/cache_config"
)

// CacheConfig - It's the custom struct used to store application configuration over the default struct defined in pkg/cache_config
// Every fields need to be declared with tags `config:""`
type CacheConfig struct {
	cacheconf.DefaultCacheConfig
	UserTimeToRefresh int `config:"USER_TIME_TO_REFRESH"`
}

var (
	cacheConfig *CacheConfig
	once        sync.Once
)

// GetInstance - Return the one struct for application configuration
func GetInstance() *CacheConfig {
	once.Do(func() {
		cacheConfig = &CacheConfig{}
		cacheconf.LoadEnvConfig(cacheConfig)
	})
	return cacheConfig
}

// GetDefaultCacheConfig - Return the instance of CacheConfigInterface
func (c *CacheConfig) GetDefaultCacheConfig() cacheconf.CacheConfigInterface {
	return &(c.DefaultCacheConfig)
}
