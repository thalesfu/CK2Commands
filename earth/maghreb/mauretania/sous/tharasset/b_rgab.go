package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉格布RgabBarony struct {
	feud.BaseBarony
}

var BRgab拉格布 feud.Barony = &拉格布RgabBarony{}

func init() {
	f := BRgab拉格布.(*拉格布RgabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rgab",
		TitleName: "拉格布",
		TitleCode: "b_rgab",
	}
}
