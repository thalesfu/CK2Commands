package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特堡OsterburgBarony struct {
	feud.BaseBarony
}

var BOsterburg奥斯特堡 feud.Barony = &奥斯特堡OsterburgBarony{}

func init() {
	f := BOsterburg奥斯特堡.(*奥斯特堡OsterburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osterburg",
		TitleName: "奥斯特堡",
		TitleCode: "b_osterburg",
	}
}
