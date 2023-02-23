package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔BarBarony struct {
	feud.BaseBarony
}

var BBar巴尔 feud.Barony = &巴尔BarBarony{}

func init() {
	f := BBar巴尔.(*巴尔BarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bar",
		TitleName: "巴尔",
		TitleCode: "b_bar",
	}
}
