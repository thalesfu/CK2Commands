package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提格济尔特TigzirtBarony struct {
	feud.BaseBarony
}

var BTigzirt提格济尔特 feud.Barony = &提格济尔特TigzirtBarony{}

func init() {
    f := BTigzirt提格济尔特.(*提格济尔特TigzirtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigzirt",
		TitleName: "提格济尔特",
		TitleCode: "b_tigzirt",
	}
}
