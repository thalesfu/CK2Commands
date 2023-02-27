package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉费夫LafeveBarony struct {
	feud.BaseBarony
}

var BLafeve拉费夫 feud.Barony = &拉费夫LafeveBarony{}

func init() {
    f := BLafeve拉费夫.(*拉费夫LafeveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lafeve",
		TitleName: "拉费夫",
		TitleCode: "b_lafeve",
	}
}
