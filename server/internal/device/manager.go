package device

import (
	"sync"

	"iot/server/internal/types"
)

type Manager struct {
	devices map[string]*types.BaseDevice
	lock    sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		devices: map[string]*types.BaseDevice{},
		lock:    sync.Mutex{},
	}
}

func (m *Manager) AddDevice(device *types.BaseDevice) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.devices[device.UUID] = &types.BaseDevice{
		UUID:     device.UUID,
		Name:     device.Name,
		Endpoint: device.Endpoint,
		Funcs:    device.Funcs,
	}
}

func (m *Manager) GetDevices() []*types.BaseDevice {
	var devices []*types.BaseDevice

	m.lock.Lock()
	defer m.lock.Unlock()

	for _, device := range m.devices {
		devices = append(devices, &types.BaseDevice{
			UUID:     device.UUID,
			Name:     device.Name,
			Endpoint: device.Endpoint,
			Funcs:    device.Funcs,
		})
	}

	return devices
}

func (m *Manager) GetDevice(uuid string) *types.BaseDevice {
	m.lock.Lock()
	defer m.lock.Unlock()

	if device, ok := m.devices[uuid]; ok {
		return &types.BaseDevice{
			UUID:     device.UUID,
			Name:     device.Name,
			Endpoint: device.Endpoint,
			Funcs:    device.Funcs,
		}
	} else {
		return nil
	}
}
