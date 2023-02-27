package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 商佉利市集Shakhari_bazarBarony struct {
	feud.BaseBarony
}

var BShakhari_bazar商佉利市集 feud.Barony = &商佉利市集Shakhari_bazarBarony{}

func init() {
    f := BShakhari_bazar商佉利市集.(*商佉利市集Shakhari_bazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shakhari_bazar",
		TitleName: "商佉利市集",
		TitleCode: "b_shakhari_bazar",
	}
}
