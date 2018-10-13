package domain

type (
	// CfgReader reads and parses a configuration file.
	CfgReader interface {
		Read(string) (Config, error)
	}
)
