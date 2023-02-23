package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑格奥莱SangoleBarony struct {
	feud.BaseBarony
}

var BSangole桑格奥莱 feud.Barony = &桑格奥莱SangoleBarony{}

func init() {
	f := BSangole桑格奥莱.(*桑格奥莱SangoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangole",
		TitleName: "桑格奥莱",
		TitleCode: "b_sangole",
	}
}
