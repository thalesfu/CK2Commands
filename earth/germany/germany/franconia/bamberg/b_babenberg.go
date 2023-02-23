package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴本堡BabenbergBarony struct {
	feud.BaseBarony
}

var BBabenberg巴本堡 feud.Barony = &巴本堡BabenbergBarony{}

func init() {
	f := BBabenberg巴本堡.(*巴本堡BabenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babenberg",
		TitleName: "巴本堡",
		TitleCode: "b_babenberg",
	}
}
