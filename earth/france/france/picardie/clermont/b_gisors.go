package clermont

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日索尔GisorsBarony struct {
	feud.BaseBarony
}

var BGisors日索尔 feud.Barony = &日索尔GisorsBarony{}

func init() {
	f := BGisors日索尔.(*日索尔GisorsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gisors",
		TitleName: "日索尔",
		TitleCode: "b_gisors",
	}
}
