package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉图CratoBarony struct {
	feud.BaseBarony
}

var BCrato克拉图 feud.Barony = &克拉图CratoBarony{}

func init() {
	f := BCrato克拉图.(*克拉图CratoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crato",
		TitleName: "克拉图",
		TitleCode: "b_crato",
	}
}
