package infra

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

type (
	// CfgReader implements domain.CfgReader .
	CfgReader struct{}
)

// Read implements domain.CfgReader#Read .
func (reader CfgReader) Read(cfgPath string) (domain.Config, error) {
	cfg := domain.Config{}
	f, err := os.Open(cfgPath)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(&cfg)
	return cfg, err
}
