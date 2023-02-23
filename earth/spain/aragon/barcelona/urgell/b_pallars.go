package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕拉尔斯PallarsBarony struct {
	feud.BaseBarony
}

var BPallars帕拉尔斯 feud.Barony = &帕拉尔斯PallarsBarony{}

func init() {
	f := BPallars帕拉尔斯.(*帕拉尔斯PallarsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pallars",
		TitleName: "帕拉尔斯",
		TitleCode: "b_pallars",
	}
}
