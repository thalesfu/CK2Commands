package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格哈特BagerhatBarony struct {
	feud.BaseBarony
}

var BBagerhat巴格哈特 feud.Barony = &巴格哈特BagerhatBarony{}

func init() {
	f := BBagerhat巴格哈特.(*巴格哈特BagerhatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagerhat",
		TitleName: "巴格哈特",
		TitleCode: "b_bagerhat",
	}
}
