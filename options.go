package yrs

// DefaultOptions returns the default options for a new Doc.
func DefaultOptions() Options {
	return Options{
		ClientID: generateClientID(),
		GUID:     generateGUID(),
	}
}
