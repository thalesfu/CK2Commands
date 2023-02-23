package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博克特拜BoqtybayBarony struct {
	feud.BaseBarony
}

var BBoqtybay博克特拜 feud.Barony = &博克特拜BoqtybayBarony{}

func init() {
	f := BBoqtybay博克特拜.(*博克特拜BoqtybayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boqtybay",
		TitleName: "博克特拜",
		TitleCode: "b_boqtybay",
	}
}
