package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏鲁ShaluBarony struct {
	feud.BaseBarony
}

var BShalu夏鲁 feud.Barony = &夏鲁ShaluBarony{}

func init() {
    f := BShalu夏鲁.(*夏鲁ShaluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shalu",
		TitleName: "夏鲁",
		TitleCode: "b_shalu",
	}
}
