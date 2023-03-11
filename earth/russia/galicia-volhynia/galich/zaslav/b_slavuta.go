package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯拉武塔SlavutaBarony struct {
	feud.BaseBarony
}

var BSlavuta斯拉武塔 feud.Barony = &斯拉武塔SlavutaBarony{}

func init() {
    f := BSlavuta斯拉武塔.(*斯拉武塔SlavutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slavuta",
		TitleName: "斯拉武塔",
		TitleCode: "b_slavuta",
	}
}
