package dns

import (
	"errors"
	"fmt"
	"os"
)

// RecordType is the type of DNS record
type RecordType string

const RecordTypeA RecordType = "a"
const RecordTypeAAAA RecordType = "aaaa"
const RecordTypeCAA RecordType = "caa"
const RecordTypeCNAME RecordType = "cname"
const RecordTypeMX RecordType = "mx"
const RecordTypeNS RecordType = "ns"
const RecordTypeSOA RecordType = "soa"
const RecordTypeSRV RecordType = "srv"
const RecordTypeTXT RecordType = "txt"

// Options is the options for the DNS client
type Options struct {
	APIKey string
	Domain string
	Types  []RecordType
}

// ValidateOptions validates the options
func (opts *Options) ValidateOptions() error {
	if opts.APIKey == "" {
		opts.APIKey = os.Getenv("CHAOS_KEY")
	}

	if opts.APIKey == "" {
		return errors.New("Authorization token not specified")
	}

	if opts.Domain == "" {
		return fmt.Errorf("No input specified for the API")
	}

	return nil
}
