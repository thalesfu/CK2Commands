package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔拉DoraBarony struct {
	feud.BaseBarony
}

var BDora达尔拉 feud.Barony = &达尔拉DoraBarony{}

func init() {
    f := BDora达尔拉.(*达尔拉DoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dora",
		TitleName: "达尔拉",
		TitleCode: "b_dora",
	}
}
