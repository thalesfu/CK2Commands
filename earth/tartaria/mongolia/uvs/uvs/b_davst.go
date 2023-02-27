package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达布斯特DavstBarony struct {
	feud.BaseBarony
}

var BDavst达布斯特 feud.Barony = &达布斯特DavstBarony{}

func init() {
    f := BDavst达布斯特.(*达布斯特DavstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "davst",
		TitleName: "达布斯特",
		TitleCode: "b_davst",
	}
}
