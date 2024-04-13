package yrs

import (
	"encoding/json"
)

// ToJSON serializes the Doc to JSON.
func (d *Doc) ToJSON() (string, error) {
	// This is a simplified example. You'll need to implement serialization for each type.
	data := map[string]interface{}{
		"clientID": d.clientID,
		"guid":     d.guid,
		// Add other fields as necessary
	}
	bytes, err := json.Marshal(data)
	return string(bytes), err
}
