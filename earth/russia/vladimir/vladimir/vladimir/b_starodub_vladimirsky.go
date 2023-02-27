package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔罗杜布弗拉基米尔斯基Starodub_vladimirskyBarony struct {
	feud.BaseBarony
}

var BStarodub_vladimirsky斯塔罗杜布弗拉基米尔斯基 feud.Barony = &斯塔罗杜布弗拉基米尔斯基Starodub_vladimirskyBarony{}

func init() {
    f := BStarodub_vladimirsky斯塔罗杜布弗拉基米尔斯基.(*斯塔罗杜布弗拉基米尔斯基Starodub_vladimirskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "starodub_vladimirsky",
		TitleName: "斯塔罗杜布弗拉基米尔斯基",
		TitleCode: "b_starodub_vladimirsky",
	}
}
