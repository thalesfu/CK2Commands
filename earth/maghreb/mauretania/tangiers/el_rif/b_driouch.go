package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代尔尤什DriouchBarony struct {
	feud.BaseBarony
}

var BDriouch代尔尤什 feud.Barony = &代尔尤什DriouchBarony{}

func init() {
    f := BDriouch代尔尤什.(*代尔尤什DriouchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "driouch",
		TitleName: "代尔尤什",
		TitleCode: "b_driouch",
	}
}
