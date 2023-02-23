package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利波他SripathaBarony struct {
	feud.BaseBarony
}

var BSripatha室利波他 feud.Barony = &室利波他SripathaBarony{}

func init() {
	f := BSripatha室利波他.(*室利波他SripathaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sripatha",
		TitleName: "室利波他",
		TitleCode: "b_sripatha",
	}
}
