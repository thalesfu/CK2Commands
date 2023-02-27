package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 设拉子Shiraz_mervBarony struct {
	feud.BaseBarony
}

var BShiraz_merv设拉子 feud.Barony = &设拉子Shiraz_mervBarony{}

func init() {
    f := BShiraz_merv设拉子.(*设拉子Shiraz_mervBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shiraz_merv",
		TitleName: "设拉子",
		TitleCode: "b_shiraz_merv",
	}
}
