package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯鲁万KairouanBarony struct {
	feud.BaseBarony
}

var BKairouan凯鲁万 feud.Barony = &凯鲁万KairouanBarony{}

func init() {
	f := BKairouan凯鲁万.(*凯鲁万KairouanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kairouan",
		TitleName: "凯鲁万",
		TitleCode: "b_kairouan",
	}
}
