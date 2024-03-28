package ws

import (
	"sync"

	"github.com/charmbracelet/log"
)

type ManagerCache struct {
	cache map[string]*Manager
	sync.RWMutex
}

func NewManagerCache() *ManagerCache {
	return &ManagerCache{
		cache: make(map[string]*Manager),
	}
}

func (mc *ManagerCache) GetManager(name string) *Manager {
	mc.RLock()
	defer mc.RUnlock()
	manager := mc.cache[name] //making sure database is the single source of truth 
	return manager
}

func (mc *ManagerCache) SetManager(name string, manager *Manager) {
	mc.Lock()
	defer mc.Unlock()
	log.Info("Creating a New Manager on Server's RAM")
	mc.cache[name] = manager
}

func (mc *ManagerCache) RemoveManager(name string) {
	mc.Lock()
	defer mc.Unlock()
	delete(mc.cache, name)
}

var MasterManager *ManagerCache = NewManagerCache()
