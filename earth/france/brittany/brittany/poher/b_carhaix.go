package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡赖CarhaixBarony struct {
	feud.BaseBarony
}

var BCarhaix卡赖 feud.Barony = &卡赖CarhaixBarony{}

func init() {
	f := BCarhaix卡赖.(*卡赖CarhaixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carhaix",
		TitleName: "卡赖",
		TitleCode: "b_carhaix",
	}
}
