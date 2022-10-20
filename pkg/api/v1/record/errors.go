package record

import "errors"

var (
	// ErrorNoRecordName is when a request / record doesn't have a primary key
	ErrorNoRecordName = errors.New("no record name")
	// ErrorNoRecordType is when a request / record doesn't have a primary key
	ErrorNoRecordType = errors.New("no record type")
	// ErrorUnsupportedType when a request for an unsupported record type occurs
	ErrorUnsupportedType = errors.New("unsupported record type")
)
