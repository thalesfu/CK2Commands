package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉纳肖奈Raith_na_seanadhBarony struct {
	feud.BaseBarony
}

var BRaith_na_seanadh拉纳肖奈 feud.Barony = &拉纳肖奈Raith_na_seanadhBarony{}

func init() {
    f := BRaith_na_seanadh拉纳肖奈.(*拉纳肖奈Raith_na_seanadhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raith_na_seanadh",
		TitleName: "拉纳肖奈",
		TitleCode: "b_raith_na_seanadh",
	}
}
