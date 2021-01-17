//go:generate statik -p assets -src=/opt/works/zplat-front/dist -dest ..
package assets

import (
	"log"
	"net/http"
	"os"

	"github.com/rakyll/statik/fs"
)

type FileSystem struct {
	efs http.FileSystem
}

func NewFS() *FileSystem {
	efs, err := fs.New()
	if err != nil {
		log.Fatalln(err)
	}

	return &FileSystem{efs}
}

func (fs FileSystem) Open(name string) (http.File, error) {
	f, err := fs.efs.Open(name)
	if os.IsNotExist(err) {
		return fs.efs.Open("/index.html") // SPA应用需要始终加载index.html
	}

	return f, err
}
