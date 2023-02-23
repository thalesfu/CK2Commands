package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布什拉BochraBarony struct {
	feud.BaseBarony
}

var BBochra布什拉 feud.Barony = &布什拉BochraBarony{}

func init() {
	f := BBochra布什拉.(*布什拉BochraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bochra",
		TitleName: "布什拉",
		TitleCode: "b_bochra",
	}
}
