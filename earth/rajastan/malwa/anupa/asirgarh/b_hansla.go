package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉罗HanslaBarony struct {
	feud.BaseBarony
}

var BHansla汉罗 feud.Barony = &汉罗HanslaBarony{}

func init() {
    f := BHansla汉罗.(*汉罗HanslaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hansla",
		TitleName: "汉罗",
		TitleCode: "b_hansla",
	}
}
