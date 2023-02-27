package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特鲁韦克OstrowekBarony struct {
	feud.BaseBarony
}

var BOstrowek奥斯特鲁韦克 feud.Barony = &奥斯特鲁韦克OstrowekBarony{}

func init() {
    f := BOstrowek奥斯特鲁韦克.(*奥斯特鲁韦克OstrowekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostrowek",
		TitleName: "奥斯特鲁韦克",
		TitleCode: "b_ostrowek",
	}
}
