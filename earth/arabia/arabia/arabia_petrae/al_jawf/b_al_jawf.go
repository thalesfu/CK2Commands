package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 焦夫Al_jawfBarony struct {
	feud.BaseBarony
}

var BAl_jawf焦夫 feud.Barony = &焦夫Al_jawfBarony{}

func init() {
    f := BAl_jawf焦夫.(*焦夫Al_jawfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_jawf",
		TitleName: "焦夫",
		TitleCode: "b_al_jawf",
	}
}
