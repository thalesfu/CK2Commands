package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎通ZatonBarony struct {
	feud.BaseBarony
}

var BZaton扎通 feud.Barony = &扎通ZatonBarony{}

func init() {
    f := BZaton扎通.(*扎通ZatonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaton",
		TitleName: "扎通",
		TitleCode: "b_zaton",
	}
}
