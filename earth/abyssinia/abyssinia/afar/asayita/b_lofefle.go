package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛菲弗勒LofefleBarony struct {
	feud.BaseBarony
}

var BLofefle洛菲弗勒 feud.Barony = &洛菲弗勒LofefleBarony{}

func init() {
    f := BLofefle洛菲弗勒.(*洛菲弗勒LofefleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lofefle",
		TitleName: "洛菲弗勒",
		TitleCode: "b_lofefle",
	}
}
