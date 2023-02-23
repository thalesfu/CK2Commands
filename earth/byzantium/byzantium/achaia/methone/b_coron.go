package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科龙CoronBarony struct {
	feud.BaseBarony
}

var BCoron科龙 feud.Barony = &科龙CoronBarony{}

func init() {
	f := BCoron科龙.(*科龙CoronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coron",
		TitleName: "科龙",
		TitleCode: "b_coron",
	}
}
