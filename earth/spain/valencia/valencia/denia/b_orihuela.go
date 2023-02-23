package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥里韦拉OrihuelaBarony struct {
	feud.BaseBarony
}

var BOrihuela奥里韦拉 feud.Barony = &奥里韦拉OrihuelaBarony{}

func init() {
	f := BOrihuela奥里韦拉.(*奥里韦拉OrihuelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orihuela",
		TitleName: "奥里韦拉",
		TitleCode: "b_orihuela",
	}
}
