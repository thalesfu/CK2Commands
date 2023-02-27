package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德拉AdraBarony struct {
	feud.BaseBarony
}

var BAdra阿德拉 feud.Barony = &阿德拉AdraBarony{}

func init() {
    f := BAdra阿德拉.(*阿德拉AdraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adra",
		TitleName: "阿德拉",
		TitleCode: "b_adra",
	}
}
