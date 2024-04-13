// NewDoc creates a new document with default options.
func NewDoc() *Doc {
	return &Doc{
		store:     &StoreRef{},
		clientID:  generateClientID(),
		guid:      generateGUID(),
		options:   DefaultOptions(),
		observers: make(map[string]interface{}),
	}
}

package yrs

import (
	"sync"
)

// Doc represents a YATA document, managing collaborative data.
type Doc struct {
	store     *StoreRef
	clientID  uint32
	guid      string
	options   Options
	observers map[string]interface{} // Placeholder for observer functions
}

// StoreRef represents a reference to the document's store.
type StoreRef struct {
	mu sync.RWMutex
	// other fields as needed
}

// Options contains configuration options for a Doc.
type Options struct {
	ClientID uint32
	GUID     string
	// other options as needed
}

// GetOrInsertText retrieves or creates a Text structure by name.
func (d *Doc) GetOrInsertText(name string) *TextRef {
	// Implementation
	return nil
}

// ObserveUpdate registers a callback for update events.
func (d *Doc) ObserveUpdate(callback func(*Transaction, *UpdateEvent)) {
	// Implementation
}
