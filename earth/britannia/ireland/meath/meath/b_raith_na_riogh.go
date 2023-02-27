package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉纳里Raith_na_rioghBarony struct {
	feud.BaseBarony
}

var BRaith_na_riogh拉纳里 feud.Barony = &拉纳里Raith_na_rioghBarony{}

func init() {
    f := BRaith_na_riogh拉纳里.(*拉纳里Raith_na_rioghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raith_na_riogh",
		TitleName: "拉纳里",
		TitleCode: "b_raith_na_riogh",
	}
}
