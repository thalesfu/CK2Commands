package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉卜SarabBarony struct {
	feud.BaseBarony
}

var BSarab萨拉卜 feud.Barony = &萨拉卜SarabBarony{}

func init() {
	f := BSarab萨拉卜.(*萨拉卜SarabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarab",
		TitleName: "萨拉卜",
		TitleCode: "b_sarab",
	}
}
