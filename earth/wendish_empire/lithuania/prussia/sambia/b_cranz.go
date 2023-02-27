package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克兰茨CranzBarony struct {
	feud.BaseBarony
}

var BCranz克兰茨 feud.Barony = &克兰茨CranzBarony{}

func init() {
    f := BCranz克兰茨.(*克兰茨CranzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cranz",
		TitleName: "克兰茨",
		TitleCode: "b_cranz",
	}
}
