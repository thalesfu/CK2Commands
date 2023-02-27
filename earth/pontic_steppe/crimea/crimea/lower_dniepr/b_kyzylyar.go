package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒亚尔KyzylyarBarony struct {
	feud.BaseBarony
}

var BKyzylyar克孜勒亚尔 feud.Barony = &克孜勒亚尔KyzylyarBarony{}

func init() {
    f := BKyzylyar克孜勒亚尔.(*克孜勒亚尔KyzylyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzylyar",
		TitleName: "克孜勒亚尔",
		TitleCode: "b_kyzylyar",
	}
}
