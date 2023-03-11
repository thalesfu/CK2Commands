package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨科夫尼纳SakovninaBarony struct {
	feud.BaseBarony
}

var BSakovnina萨科夫尼纳 feud.Barony = &萨科夫尼纳SakovninaBarony{}

func init() {
    f := BSakovnina萨科夫尼纳.(*萨科夫尼纳SakovninaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakovnina",
		TitleName: "萨科夫尼纳",
		TitleCode: "b_sakovnina",
	}
}
