package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯克RasakBarony struct {
	feud.BaseBarony
}

var BRasak拉斯克 feud.Barony = &拉斯克RasakBarony{}

func init() {
	f := BRasak拉斯克.(*拉斯克RasakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rasak",
		TitleName: "拉斯克",
		TitleCode: "b_rasak",
	}
}
