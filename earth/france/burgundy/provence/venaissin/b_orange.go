package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥朗热OrangeBarony struct {
	feud.BaseBarony
}

var BOrange奥朗热 feud.Barony = &奥朗热OrangeBarony{}

func init() {
	f := BOrange奥朗热.(*奥朗热OrangeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orange",
		TitleName: "奥朗热",
		TitleCode: "b_orange",
	}
}
