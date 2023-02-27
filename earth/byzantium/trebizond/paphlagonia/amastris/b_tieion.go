package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提翁TieionBarony struct {
	feud.BaseBarony
}

var BTieion提翁 feud.Barony = &提翁TieionBarony{}

func init() {
    f := BTieion提翁.(*提翁TieionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tieion",
		TitleName: "提翁",
		TitleCode: "b_tieion",
	}
}
