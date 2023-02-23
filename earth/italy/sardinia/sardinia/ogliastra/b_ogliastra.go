package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利亚斯特拉OgliastraBarony struct {
	feud.BaseBarony
}

var BOgliastra奥利亚斯特拉 feud.Barony = &奥利亚斯特拉OgliastraBarony{}

func init() {
	f := BOgliastra奥利亚斯特拉.(*奥利亚斯特拉OgliastraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ogliastra",
		TitleName: "奥利亚斯特拉",
		TitleCode: "b_ogliastra",
	}
}
