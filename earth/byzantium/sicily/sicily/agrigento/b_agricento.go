package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔真蒂AgricentoBarony struct {
	feud.BaseBarony
}

var BAgricento吉尔真蒂 feud.Barony = &吉尔真蒂AgricentoBarony{}

func init() {
    f := BAgricento吉尔真蒂.(*吉尔真蒂AgricentoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agricento",
		TitleName: "吉尔真蒂",
		TitleCode: "b_agricento",
	}
}
