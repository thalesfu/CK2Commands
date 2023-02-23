package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弥室利迦MisrikhBarony struct {
	feud.BaseBarony
}

var BMisrikh弥室利迦 feud.Barony = &弥室利迦MisrikhBarony{}

func init() {
	f := BMisrikh弥室利迦.(*弥室利迦MisrikhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "misrikh",
		TitleName: "弥室利迦",
		TitleCode: "b_misrikh",
	}
}
