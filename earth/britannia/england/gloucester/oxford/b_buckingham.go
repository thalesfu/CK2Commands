package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白金汉BuckinghamBarony struct {
	feud.BaseBarony
}

var BBuckingham白金汉 feud.Barony = &白金汉BuckinghamBarony{}

func init() {
    f := BBuckingham白金汉.(*白金汉BuckinghamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buckingham",
		TitleName: "白金汉",
		TitleCode: "b_buckingham",
	}
}
