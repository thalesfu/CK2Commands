package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索伦托SorrentoBarony struct {
	feud.BaseBarony
}

var BSorrento索伦托 feud.Barony = &索伦托SorrentoBarony{}

func init() {
	f := BSorrento索伦托.(*索伦托SorrentoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorrento",
		TitleName: "索伦托",
		TitleCode: "b_sorrento",
	}
}
