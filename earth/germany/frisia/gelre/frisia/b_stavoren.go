package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔福伦StavorenBarony struct {
	feud.BaseBarony
}

var BStavoren斯塔福伦 feud.Barony = &斯塔福伦StavorenBarony{}

func init() {
	f := BStavoren斯塔福伦.(*斯塔福伦StavorenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stavoren",
		TitleName: "斯塔福伦",
		TitleCode: "b_stavoren",
	}
}
