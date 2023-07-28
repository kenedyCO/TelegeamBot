package files

import (
	"github.com/kenedyCO/tgBot/lib/e"
)

type Storage struct {
	k
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *Storage.page) (err error) {
	defer func() { err = e.WrapIfErr("can't get updates", err) }()

}
