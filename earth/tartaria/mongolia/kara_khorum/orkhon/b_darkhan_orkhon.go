package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔汗Darkhan_orkhonBarony struct {
	feud.BaseBarony
}

var BDarkhan_orkhon达尔汗 feud.Barony = &达尔汗Darkhan_orkhonBarony{}

func init() {
    f := BDarkhan_orkhon达尔汗.(*达尔汗Darkhan_orkhonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darkhan_orkhon",
		TitleName: "达尔汗",
		TitleCode: "b_darkhan_orkhon",
	}
}
