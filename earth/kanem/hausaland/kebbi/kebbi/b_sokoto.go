package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索科托SokotoBarony struct {
	feud.BaseBarony
}

var BSokoto索科托 feud.Barony = &索科托SokotoBarony{}

func init() {
	f := BSokoto索科托.(*索科托SokotoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sokoto",
		TitleName: "索科托",
		TitleCode: "b_sokoto",
	}
}
