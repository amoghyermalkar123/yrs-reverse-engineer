package yrs

// Transaction represents a transaction on the document.
type Transaction struct {
	// fields as needed
}

// TransactionMut represents a mutable transaction on the document.
type TransactionMut struct {
	// fields as needed
}

// NewTransaction creates a new read-only transaction.
func (d *Doc) NewTransaction() *Transaction {
	// Implementation details here
	return &Transaction{}
}

// NewTransactionMut creates a new read-write transaction.
func (d *Doc) NewTransactionMut() *TransactionMut {
	// Implementation details here
	return &TransactionMut{}
}
