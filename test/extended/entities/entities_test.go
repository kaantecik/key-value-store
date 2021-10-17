package entities

import (
	"github.com/kaantecik/key-value-store/internal/config"
	"github.com/kaantecik/key-value-store/internal/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheOptions(t *testing.T) {
	options := &entities.CacheOptions{}
	_ = entities.NewCache(options)
	assert.Equal(t, config.DefaultSaveLocation, options.SaveLocation)
	assert.Equal(t, config.DefaultSaveInterval, options.SaveInterval)
}

func TestGetSetItem(t *testing.T) {
	var (
		testKey   = "foo"
		testValue = "bar"
	)

	options := &entities.CacheOptions{}
	c := entities.NewCache(options)
	c.Set(testKey, testValue)

	item, found := c.Get(testKey)

	assert.Equal(t, true, found)
	assert.Equal(t, testKey, item.Key)
	assert.Equal(t, testValue, item.Value)
}

func TestFlush(t *testing.T) {
	var (
		testKey   = "foo"
		testValue = "bar"
	)
	options := &entities.CacheOptions{}
	c := entities.NewCache(options)
	c.Set(testKey, testValue)

	c.Flush()

	_, found := c.Get(testKey)

	assert.Equal(t, false, found)
}
