package entities

import (
	"encoding/json"
	"fmt"
	"github.com/kaantecik/key-value-store/internal/config"
	"github.com/kaantecik/key-value-store/internal/logging"
	"github.com/kaantecik/key-value-store/tools/io"
	"os"
	"strings"
	"sync"
	"time"
)

// Cache is main part of app.
type Cache struct {
	id      string
	items   map[string]*Item
	mutex   sync.RWMutex
	options *CacheOptions
}

// CacheOptions defines the options for Cache.
type CacheOptions struct {
	// SaveInterval represents interval of saving data.
	//
	// Default: 10 * time.Minute
	SaveInterval time.Duration

	// SaveLocation represents where the log files saved.
	// All log files will be saved in this format: TIMESTAMP-data.json
	// Default: /tmp/kv-store/logs
	SaveLocation string
}

type cacheFile struct {
	Items []*Item `json:"items"`
}

// setInterval function sets new interval to Cache
func (c *Cache) setInterval() chan<- bool {
	ticker := time.NewTicker(c.options.SaveInterval)
	stopChan := make(chan bool)
	go func() {
		for {
			select {
			case <-stopChan:
				return
			case <-ticker.C:
				c.saveFiles()
			}
		}
	}()

	return stopChan
}

// saveFiles function save files on given interval.
func (c *Cache) saveFiles() {
	logging.AppLogger.Info(config.MessageSavingFiles)

	path := fmt.Sprintf("%s/%v-data.json", config.DefaultSaveLocation, time.Now().Unix())

	f, err := os.Create(path)

	if err != nil {
		logging.ErrorLogger.Fatal(err)
	}

	for _, item := range c.items {
		conv, err := json.Marshal(item)
		if err != nil {
			logging.ErrorLogger.Fatal(err)
		}
		_, err = f.Write(conv)
		if err != nil {
			logging.ErrorLogger.Fatal(err)
		}
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logging.ErrorLogger.Fatal(err)
		}
	}(f)
}

// Get function returns an item from the cache. Returns the item if exists.
func (c *Cache) Get(key string) (*Item, bool) {
	c.mutex.RLock()

	item, found := c.items[key]

	if !found {
		c.mutex.RUnlock()
		return nil, false
	}

	c.mutex.RUnlock()

	return item, true
}

// SetItem is a function that create new Item. Accepts *Item
func (c *Cache) SetItem(item *Item) {
	c.mutex.RLock()

	logging.AppLogger.Infof("| %s | %v |", config.MessageCreated, item)

	c.items[item.Key] = item

	c.mutex.RUnlock()
}

// Set is a function that create new Item.
// Accepts key and value.
func (c *Cache) Set(key string, value interface{}) {
	c.mutex.RLock()

	logging.AppLogger.Infof("| %s | %s : %v |", config.MessageCreated, key, value)

	c.items[key] = NewItem(key, value)

	c.mutex.RUnlock()
}

// Flush is a function that flushes Cache.
func (c *Cache) Flush() {
	c.mutex.RLock()
	c.items = map[string]*Item{}
	logging.AppLogger.Infof("| %s | %v |", config.MessageCacheFlushed, c.items)
	c.mutex.RUnlock()
}

// NewCache is a factory for creating new cache.
func NewCache(options *CacheOptions) *Cache {
	// Set default internal if DefaultSaveInterval does not exist in options.
	if options.SaveInterval == 0*time.Second {
		options.SaveInterval = config.DefaultSaveInterval
	}

	// Set default location if DefaultSaveLocation does not exist in options.
	if options.SaveLocation == "" {
		options.SaveLocation = config.DefaultSaveLocation
	}

	cache := &Cache{
		items:   map[string]*Item{},
		options: options,
	}

	cache.setInterval()

	if _, err := os.Stat(config.DefaultSaveLocation); os.IsNotExist(err) {
		iotools.CreateFolder(config.DefaultSaveLocation)
	}

	// Create if file does not exist
	if _, err := os.Stat(config.LogPath); os.IsNotExist(err) {
		iotools.CreateFolder(config.LogPath)
	} else {
		files := iotools.GetFiles(options.SaveLocation)

		for _, file := range files {
			if strings.Contains(file, config.AllowedConfigExt) {
				iotools.ReadFileAndProcess(file, func(param []byte) {
					var file cacheFile

					err := json.Unmarshal(param, &file)

					if err != nil {
						logging.ErrorLogger.Error(err)
					}

					for _, item := range file.Items {
						cache.SetItem(item)
						logging.AppLogger.Info(item)
					}
				})
			}
		}
	}

	return cache
}
