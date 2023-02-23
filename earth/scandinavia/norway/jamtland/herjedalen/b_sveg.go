package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯韦格SvegBarony struct {
	feud.BaseBarony
}

var BSveg斯韦格 feud.Barony = &斯韦格SvegBarony{}

func init() {
	f := BSveg斯韦格.(*斯韦格SvegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sveg",
		TitleName: "斯韦格",
		TitleCode: "b_sveg",
	}
}
