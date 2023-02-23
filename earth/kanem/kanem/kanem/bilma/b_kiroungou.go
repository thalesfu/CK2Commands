package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基伦古KiroungouBarony struct {
	feud.BaseBarony
}

var BKiroungou基伦古 feud.Barony = &基伦古KiroungouBarony{}

func init() {
	f := BKiroungou基伦古.(*基伦古KiroungouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiroungou",
		TitleName: "基伦古",
		TitleCode: "b_kiroungou",
	}
}
