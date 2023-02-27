package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩谛MaltherBarony struct {
	feud.BaseBarony
}

var BMalther摩谛 feud.Barony = &摩谛MaltherBarony{}

func init() {
    f := BMalther摩谛.(*摩谛MaltherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malther",
		TitleName: "摩谛",
		TitleCode: "b_malther",
	}
}
