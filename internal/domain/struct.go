package domain

type (
	// Config represents configuration.
	Config struct {
		Default DefaultConfiguration
		Items   []Item
	}

	// DefaultConfiguration represents default configuration.
	DefaultConfiguration struct {
		SrcDefaultConfiguration
	}

	// SrcDefaultConfiguration represents source's default configuration.
	SrcDefaultConfiguration struct {
		InterfacePrefix string `yaml:"interface_prefix"`
		InterfaceSuffix string `yaml:"interface_suffix"`
		VendorDir       string `yaml:"vendor_dir"`
	}

	// Item represents configuration of each mock.
	Item struct {
		Src  Src
		Dest Dest
		DefaultConfiguration
	}

	// Src represents source configuration.
	Src struct {
		Package   string
		Interface string
		Name      string
		File      string
		Dir       string
		SrcDefaultConfiguration
	}

	// Dest represents generated file's configuration.
	Dest struct {
		Package string
		File    string
	}
)
