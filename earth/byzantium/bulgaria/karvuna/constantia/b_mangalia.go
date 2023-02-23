package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼加利亚MangaliaBarony struct {
	feud.BaseBarony
}

var BMangalia曼加利亚 feud.Barony = &曼加利亚MangaliaBarony{}

func init() {
	f := BMangalia曼加利亚.(*曼加利亚MangaliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangalia",
		TitleName: "曼加利亚",
		TitleCode: "b_mangalia",
	}
}
