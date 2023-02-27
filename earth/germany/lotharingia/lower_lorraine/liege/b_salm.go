package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔姆SalmBarony struct {
	feud.BaseBarony
}

var BSalm萨尔姆 feud.Barony = &萨尔姆SalmBarony{}

func init() {
    f := BSalm萨尔姆.(*萨尔姆SalmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salm",
		TitleName: "萨尔姆",
		TitleCode: "b_salm",
	}
}
