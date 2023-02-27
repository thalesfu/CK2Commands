package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琐珥ZoarBarony struct {
	feud.BaseBarony
}

var BZoar琐珥 feud.Barony = &琐珥ZoarBarony{}

func init() {
    f := BZoar琐珥.(*琐珥ZoarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zoar",
		TitleName: "琐珥",
		TitleCode: "b_zoar",
	}
}
