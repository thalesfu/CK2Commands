package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科兹拉CozlaBarony struct {
	feud.BaseBarony
}

var BCozla科兹拉 feud.Barony = &科兹拉CozlaBarony{}

func init() {
    f := BCozla科兹拉.(*科兹拉CozlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cozla",
		TitleName: "科兹拉",
		TitleCode: "b_cozla",
	}
}
