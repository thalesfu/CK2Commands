package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡尼SignyBarony struct {
	feud.BaseBarony
}

var BSigny锡尼 feud.Barony = &锡尼SignyBarony{}

func init() {
	f := BSigny锡尼.(*锡尼SignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "signy",
		TitleName: "锡尼",
		TitleCode: "b_signy",
	}
}
