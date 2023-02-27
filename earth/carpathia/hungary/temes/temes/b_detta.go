package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代塔DettaBarony struct {
	feud.BaseBarony
}

var BDetta代塔 feud.Barony = &代塔DettaBarony{}

func init() {
    f := BDetta代塔.(*代塔DettaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "detta",
		TitleName: "代塔",
		TitleCode: "b_detta",
	}
}
