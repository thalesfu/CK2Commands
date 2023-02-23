package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肖普朗SopronBarony struct {
	feud.BaseBarony
}

var BSopron肖普朗 feud.Barony = &肖普朗SopronBarony{}

func init() {
	f := BSopron肖普朗.(*肖普朗SopronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sopron",
		TitleName: "肖普朗",
		TitleCode: "b_sopron",
	}
}
