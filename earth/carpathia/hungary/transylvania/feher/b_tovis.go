package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特维什TovisBarony struct {
	feud.BaseBarony
}

var BTovis特维什 feud.Barony = &特维什TovisBarony{}

func init() {
    f := BTovis特维什.(*特维什TovisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tovis",
		TitleName: "特维什",
		TitleCode: "b_tovis",
	}
}
