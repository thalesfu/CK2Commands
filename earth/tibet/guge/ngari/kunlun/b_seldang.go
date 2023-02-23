package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 色当SeldangBarony struct {
	feud.BaseBarony
}

var BSeldang色当 feud.Barony = &色当SeldangBarony{}

func init() {
	f := BSeldang色当.(*色当SeldangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seldang",
		TitleName: "色当",
		TitleCode: "b_seldang",
	}
}
