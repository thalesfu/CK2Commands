package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆斯马尔MusmerBarony struct {
	feud.BaseBarony
}

var BMusmer穆斯马尔 feud.Barony = &穆斯马尔MusmerBarony{}

func init() {
	f := BMusmer穆斯马尔.(*穆斯马尔MusmerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "musmer",
		TitleName: "穆斯马尔",
		TitleCode: "b_musmer",
	}
}
