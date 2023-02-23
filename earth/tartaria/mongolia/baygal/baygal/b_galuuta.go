package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嘎鲁特GaluutaBarony struct {
	feud.BaseBarony
}

var BGaluuta嘎鲁特 feud.Barony = &嘎鲁特GaluutaBarony{}

func init() {
	f := BGaluuta嘎鲁特.(*嘎鲁特GaluutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galuuta",
		TitleName: "嘎鲁特",
		TitleCode: "b_galuuta",
	}
}
