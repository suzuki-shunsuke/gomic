package gencmd

import (
	"fmt"
	"path/filepath"

	"github.com/suzuki-shunsuke/go-cliutil"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

// Main is usecase layer's entrypoint of gen command.
func Main(
	fsys domain.FileSystem, importer domain.Importer,
	cfgReader domain.CfgReader, cfgPath string,
) error {
	if cfgPath == "" {
		wd, err := fsys.Getwd()
		if err != nil {
			return err
		}
		d, err := cliutil.FindFile(wd, fsys.Exist, ".gomic.yml")
		if err != nil {
			return err
		}
		cfgPath = d
	}
	var err error
	cfgPath, err = filepath.Abs(cfgPath)
	if err != nil {
		return err
	}
	cfg, err := cfgReader.Read(cfgPath)
	if err != nil {
		return err
	}
	// TODO validation
	cfg, err = initCfg(cfg, cfgPath)
	if err != nil {
		return err
	}
	cfgDir := filepath.Dir(cfgPath)
	for _, item := range cfg.Items {
		if err := genFile(fsys, importer, item, cfgDir); err != nil {
			return err
		}
	}
	return nil
}

func initCfg(cfg domain.Config, cfgPath string) (domain.Config, error) {
	cfgDir := filepath.Dir(cfgPath)
	for i, item := range cfg.Items {
		if item.InterfacePrefix == "" {
			item.InterfacePrefix = cfg.Default.InterfacePrefix
		}
		if item.InterfaceSuffix == "" {
			item.InterfaceSuffix = cfg.Default.InterfaceSuffix
		}
		if item.VendorDir == "" {
			item.VendorDir = cfg.Default.VendorDir
		}
		if item.Src.InterfacePrefix == "" {
			item.Src.InterfacePrefix = item.InterfacePrefix
		}
		if item.Src.InterfaceSuffix == "" {
			item.Src.InterfaceSuffix = item.InterfaceSuffix
		}
		if item.Src.VendorDir == "" {
			item.Src.VendorDir = item.VendorDir
		}
		if !filepath.IsAbs(item.Dest.File) {
			item.Dest.File = filepath.Join(cfgDir, item.Dest.File)
		}
		if item.Src.Name == "" {
			item.Src.Name = fmt.Sprintf(
				"%s%s%s", item.Src.InterfacePrefix,
				item.Src.Interface, item.Src.InterfaceSuffix)
		}

		if item.Src.File != "" && !filepath.IsAbs(item.Src.File) {
			item.Src.File = filepath.Join(cfgDir, item.Src.File)
		}
		if item.Src.Dir != "" && !filepath.IsAbs(item.Src.Dir) {
			item.Src.Dir = filepath.Join(cfgDir, item.Src.Dir)
		}
		if !filepath.IsAbs(item.Src.VendorDir) {
			item.Src.VendorDir = filepath.Join(cfgDir, item.Src.VendorDir)
		}
		cfg.Items[i] = item
	}
	return cfg, nil
}

func genFile(
	fsys domain.FileSystem, importer domain.Importer,
	item domain.Item, cfgDir string,
) error {
	if item.Dest.Package == "" {
		pkg, err := importer.GetBuildPkgInDir(filepath.Dir(item.Dest.File))
		if err != nil {
			return err
		}
		item.Dest.Package = pkg.Name
	}
	if item.Src.Package != "" {
		mock, err := parsePkg(importer, item)
		if err != nil {
			return err
		}
		return renderMock(fsys, item.Dest.File, mock)
	}
	if item.Src.File != "" {
		mock, err := parseFile(importer, item, item.Src.File)
		if err != nil {
			return err
		}
		return renderMock(fsys, item.Dest.File, mock)
	}
	if item.Src.Dir != "" {
		mock, err := parseDir(importer, item, item.Src.Dir)
		if err != nil {
			return err
		}
		return renderMock(fsys, item.Dest.File, mock)
	}
	return fmt.Errorf("file or dir or package are required")
}

func renderMock(
	fsys domain.FileSystem, destFilePath string, mock domain.MockTemplateArg,
) error {
	w, err := fsys.GetWriteCloser(destFilePath)
	if err != nil {
		return err
	}
	defer w.Close()
	return renderTpl(w, domain.MockTpl, mock)
}
