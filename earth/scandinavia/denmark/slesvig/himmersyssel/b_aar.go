package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔AarBarony struct {
	feud.BaseBarony
}

var BAar奥尔 feud.Barony = &奥尔AarBarony{}

func init() {
	f := BAar奥尔.(*奥尔AarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aar",
		TitleName: "奥尔",
		TitleCode: "b_aar",
	}
}
