package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索里亚SoriaBarony struct {
	feud.BaseBarony
}

var BSoria索里亚 feud.Barony = &索里亚SoriaBarony{}

func init() {
	f := BSoria索里亚.(*索里亚SoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soria",
		TitleName: "索里亚",
		TitleCode: "b_soria",
	}
}
