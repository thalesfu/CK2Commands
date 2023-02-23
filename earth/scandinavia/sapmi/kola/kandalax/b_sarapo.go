package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉帕SarapoBarony struct {
	feud.BaseBarony
}

var BSarapo萨拉帕 feud.Barony = &萨拉帕SarapoBarony{}

func init() {
	f := BSarapo萨拉帕.(*萨拉帕SarapoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarapo",
		TitleName: "萨拉帕",
		TitleCode: "b_sarapo",
	}
}
