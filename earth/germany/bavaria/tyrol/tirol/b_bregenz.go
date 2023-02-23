package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷根茨BregenzBarony struct {
	feud.BaseBarony
}

var BBregenz布雷根茨 feud.Barony = &布雷根茨BregenzBarony{}

func init() {
	f := BBregenz布雷根茨.(*布雷根茨BregenzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bregenz",
		TitleName: "布雷根茨",
		TitleCode: "b_bregenz",
	}
}
