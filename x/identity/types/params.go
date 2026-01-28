package types

import (
	"fmt"
	"yaml"
)

// DefaultParams returns the default module parameters
func DefaultParams() Params {
	return Params{
		MaxUsernameLength: 32,
		DefaultMetadataUri: "ipfs://default",
	}
}

// Validate validates the params
func (p Params) Validate() error {
	if p.MaxUsernameLength <= 0 {
		return fmt.Errorf("max_username_length must be positive")
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}