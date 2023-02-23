package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讷韦尔NeversBarony struct {
	feud.BaseBarony
}

var BNevers讷韦尔 feud.Barony = &讷韦尔NeversBarony{}

func init() {
	f := BNevers讷韦尔.(*讷韦尔NeversBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nevers",
		TitleName: "讷韦尔",
		TitleCode: "b_nevers",
	}
}
