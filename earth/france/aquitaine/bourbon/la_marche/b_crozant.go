package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗藏CrozantBarony struct {
	feud.BaseBarony
}

var BCrozant克罗藏 feud.Barony = &克罗藏CrozantBarony{}

func init() {
    f := BCrozant克罗藏.(*克罗藏CrozantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crozant",
		TitleName: "克罗藏",
		TitleCode: "b_crozant",
	}
}
