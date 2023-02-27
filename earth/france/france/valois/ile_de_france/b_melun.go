package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默伦MelunBarony struct {
	feud.BaseBarony
}

var BMelun默伦 feud.Barony = &默伦MelunBarony{}

func init() {
    f := BMelun默伦.(*默伦MelunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melun",
		TitleName: "默伦",
		TitleCode: "b_melun",
	}
}
