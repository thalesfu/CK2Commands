package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代希利耶DakhiliyahBarony struct {
	feud.BaseBarony
}

var BDakhiliyah代希利耶 feud.Barony = &代希利耶DakhiliyahBarony{}

func init() {
    f := BDakhiliyah代希利耶.(*代希利耶DakhiliyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dakhiliyah",
		TitleName: "代希利耶",
		TitleCode: "b_dakhiliyah",
	}
}
