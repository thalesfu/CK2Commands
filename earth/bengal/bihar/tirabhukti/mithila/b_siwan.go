package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡万SiwanBarony struct {
	feud.BaseBarony
}

var BSiwan锡万 feud.Barony = &锡万SiwanBarony{}

func init() {
    f := BSiwan锡万.(*锡万SiwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siwan",
		TitleName: "锡万",
		TitleCode: "b_siwan",
	}
}
