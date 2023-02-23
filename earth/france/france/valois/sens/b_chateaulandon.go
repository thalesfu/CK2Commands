package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗东堡ChateaulandonBarony struct {
	feud.BaseBarony
}

var BChateaulandon朗东堡 feud.Barony = &朗东堡ChateaulandonBarony{}

func init() {
	f := BChateaulandon朗东堡.(*朗东堡ChateaulandonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateaulandon",
		TitleName: "朗东堡",
		TitleCode: "b_chateaulandon",
	}
}
