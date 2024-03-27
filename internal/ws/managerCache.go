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
	manager, ok := mc.cache[name]
	if !ok {
		log.Info("Creating Debug Manager")
		m := NewManager("root")
		mc.cache["root"] = m
		return m
	}
	log.Info("CACHE HIT!")
	return manager
}

func (mc *ManagerCache) SetManager(name string, manager *Manager) {
	mc.Lock()
	defer mc.Unlock()
	mc.cache[name] = manager
}

func (mc *ManagerCache) RemoveManager(name string) {
	mc.Lock()
	defer mc.Unlock()
	delete(mc.cache, name)
}

var MasterManager *ManagerCache = NewManagerCache()
