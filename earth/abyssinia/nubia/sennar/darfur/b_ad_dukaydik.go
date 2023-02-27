package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜凯迪克Ad_dukaydikBarony struct {
	feud.BaseBarony
}

var BAd_dukaydik杜凯迪克 feud.Barony = &杜凯迪克Ad_dukaydikBarony{}

func init() {
    f := BAd_dukaydik杜凯迪克.(*杜凯迪克Ad_dukaydikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ad_dukaydik",
		TitleName: "杜凯迪克",
		TitleCode: "b_ad_dukaydik",
	}
}
