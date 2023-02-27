package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒_亚尔KyzalyarBarony struct {
	feud.BaseBarony
}

var BKyzalyar克孜勒_亚尔 feud.Barony = &克孜勒_亚尔KyzalyarBarony{}

func init() {
    f := BKyzalyar克孜勒_亚尔.(*克孜勒_亚尔KyzalyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzalyar",
		TitleName: "克孜勒_亚尔",
		TitleCode: "b_kyzalyar",
	}
}
