package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 畔荼娑罗BhandasarBarony struct {
	feud.BaseBarony
}

var BBhandasar畔荼娑罗 feud.Barony = &畔荼娑罗BhandasarBarony{}

func init() {
	f := BBhandasar畔荼娑罗.(*畔荼娑罗BhandasarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhandasar",
		TitleName: "畔荼娑罗",
		TitleCode: "b_bhandasar",
	}
}
