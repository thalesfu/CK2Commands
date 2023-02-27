package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯蒂亚OstiaBarony struct {
	feud.BaseBarony
}

var BOstia奥斯蒂亚 feud.Barony = &奥斯蒂亚OstiaBarony{}

func init() {
    f := BOstia奥斯蒂亚.(*奥斯蒂亚OstiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostia",
		TitleName: "奥斯蒂亚",
		TitleCode: "b_ostia",
	}
}
