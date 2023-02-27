package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔德克SaridagBarony struct {
	feud.BaseBarony
}

var BSaridag萨尔德克 feud.Barony = &萨尔德克SaridagBarony{}

func init() {
    f := BSaridag萨尔德克.(*萨尔德克SaridagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saridag",
		TitleName: "萨尔德克",
		TitleCode: "b_saridag",
	}
}
