package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米利都MiletosBarony struct {
	feud.BaseBarony
}

var BMiletos米利都 feud.Barony = &米利都MiletosBarony{}

func init() {
    f := BMiletos米利都.(*米利都MiletosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miletos",
		TitleName: "米利都",
		TitleCode: "b_miletos",
	}
}
