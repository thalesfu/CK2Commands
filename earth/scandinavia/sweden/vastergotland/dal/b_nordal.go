package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺德尔NordalBarony struct {
	feud.BaseBarony
}

var BNordal诺德尔 feud.Barony = &诺德尔NordalBarony{}

func init() {
	f := BNordal诺德尔.(*诺德尔NordalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nordal",
		TitleName: "诺德尔",
		TitleCode: "b_nordal",
	}
}
