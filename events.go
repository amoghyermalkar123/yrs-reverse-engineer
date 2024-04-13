package yrs

import "time"

// UpdateEvent represents an event triggered by a document update.
type UpdateEvent struct {
	Transaction *TransactionMut // The transaction that triggered the update event.
	Data        []byte          // Encoded data representing the changes made in the transaction.
	Timestamp   time.Time       // Timestamp when the event was triggered.
	Origin      string          // Origin of the update, e.g., "local" or "remote".
}

// NewUpdateEvent creates a new instance of UpdateEvent.
func NewUpdateEvent(txn *TransactionMut, data []byte, origin string) *UpdateEvent {
	return &UpdateEvent{
		Transaction: txn,
		Data:        data,
		Timestamp:   time.Now(),
		Origin:      origin,
	}
}
