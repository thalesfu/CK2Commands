package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法希FachiBarony struct {
	feud.BaseBarony
}

var BFachi法希 feud.Barony = &法希FachiBarony{}

func init() {
	f := BFachi法希.(*法希FachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fachi",
		TitleName: "法希",
		TitleCode: "b_fachi",
	}
}
