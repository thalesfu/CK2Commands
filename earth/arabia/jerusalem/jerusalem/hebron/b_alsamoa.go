package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾萨摩亚AlsamoaBarony struct {
	feud.BaseBarony
}

var BAlsamoa艾萨摩亚 feud.Barony = &艾萨摩亚AlsamoaBarony{}

func init() {
    f := BAlsamoa艾萨摩亚.(*艾萨摩亚AlsamoaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alsamoa",
		TitleName: "艾萨摩亚",
		TitleCode: "b_alsamoa",
	}
}
