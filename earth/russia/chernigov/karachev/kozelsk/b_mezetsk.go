package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅泽茨克MezetskBarony struct {
	feud.BaseBarony
}

var BMezetsk梅泽茨克 feud.Barony = &梅泽茨克MezetskBarony{}

func init() {
	f := BMezetsk梅泽茨克.(*梅泽茨克MezetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mezetsk",
		TitleName: "梅泽茨克",
		TitleCode: "b_mezetsk",
	}
}
