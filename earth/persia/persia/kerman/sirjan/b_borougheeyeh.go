package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博鲁盖耶BorougheeyehBarony struct {
	feud.BaseBarony
}

var BBorougheeyeh博鲁盖耶 feud.Barony = &博鲁盖耶BorougheeyehBarony{}

func init() {
	f := BBorougheeyeh博鲁盖耶.(*博鲁盖耶BorougheeyehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borougheeyeh",
		TitleName: "博鲁盖耶",
		TitleCode: "b_borougheeyeh",
	}
}
