package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔普TerpeBarony struct {
	feud.BaseBarony
}

var BTerpe特尔普 feud.Barony = &特尔普TerpeBarony{}

func init() {
    f := BTerpe特尔普.(*特尔普TerpeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terpe",
		TitleName: "特尔普",
		TitleCode: "b_terpe",
	}
}
