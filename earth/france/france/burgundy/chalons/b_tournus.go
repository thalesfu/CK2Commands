package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔尼TournusBarony struct {
	feud.BaseBarony
}

var BTournus图尔尼 feud.Barony = &图尔尼TournusBarony{}

func init() {
	f := BTournus图尔尼.(*图尔尼TournusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tournus",
		TitleName: "图尔尼",
		TitleCode: "b_tournus",
	}
}
