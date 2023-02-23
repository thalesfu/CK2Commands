package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔MerleBarony struct {
	feud.BaseBarony
}

var BMerle梅尔 feud.Barony = &梅尔MerleBarony{}

func init() {
	f := BMerle梅尔.(*梅尔MerleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merle",
		TitleName: "梅尔",
		TitleCode: "b_merle",
	}
}
