package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 申SkonBarony struct {
	feud.BaseBarony
}

var BSkon申 feud.Barony = &申SkonBarony{}

func init() {
	f := BSkon申.(*申SkonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skon",
		TitleName: "申",
		TitleCode: "b_skon",
	}
}
