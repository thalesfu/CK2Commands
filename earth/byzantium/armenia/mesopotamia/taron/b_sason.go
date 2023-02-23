package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨松SasonBarony struct {
	feud.BaseBarony
}

var BSason萨松 feud.Barony = &萨松SasonBarony{}

func init() {
	f := BSason萨松.(*萨松SasonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sason",
		TitleName: "萨松",
		TitleCode: "b_sason",
	}
}
