package clermont

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢勒ChellesBarony struct {
	feud.BaseBarony
}

var BChelles谢勒 feud.Barony = &谢勒ChellesBarony{}

func init() {
	f := BChelles谢勒.(*谢勒ChellesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chelles",
		TitleName: "谢勒",
		TitleCode: "b_chelles",
	}
}
