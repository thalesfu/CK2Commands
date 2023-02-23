package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒沃克ShowakBarony struct {
	feud.BaseBarony
}

var BShowak舒沃克 feud.Barony = &舒沃克ShowakBarony{}

func init() {
	f := BShowak舒沃克.(*舒沃克ShowakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "showak",
		TitleName: "舒沃克",
		TitleCode: "b_showak",
	}
}
