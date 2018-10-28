package initcmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

// Main is usecase layer's entrypoint of init command.
func Main(fsys domain.FileSystem, dst string) error {
	if fsys.Exist(dst) {
		return nil
	}
	dir := filepath.Dir(dst)
	if !fsys.Exist(dir) {
		fmt.Printf("create a directory %s\n", dir)
		if err := fsys.MkdirAll(dir); err != nil {
			return err
		}
	}
	fmt.Printf("create %s\n", dst)
	return fsys.Write(dst, []byte(strings.Trim(domain.ConfigTpl, "\n")))
}
