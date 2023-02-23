package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉桑根RathanganBarony struct {
	feud.BaseBarony
}

var BRathangan拉桑根 feud.Barony = &拉桑根RathanganBarony{}

func init() {
	f := BRathangan拉桑根.(*拉桑根RathanganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rathangan",
		TitleName: "拉桑根",
		TitleCode: "b_rathangan",
	}
}
