package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法里曼FarimBarony struct {
	feud.BaseBarony
}

var BFarim法里曼 feud.Barony = &法里曼FarimBarony{}

func init() {
	f := BFarim法里曼.(*法里曼FarimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farim",
		TitleName: "法里曼",
		TitleCode: "b_farim",
	}
}
