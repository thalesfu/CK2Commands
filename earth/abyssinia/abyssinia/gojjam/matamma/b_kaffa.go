package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 咖法KaffaBarony struct {
	feud.BaseBarony
}

var BKaffa咖法 feud.Barony = &咖法KaffaBarony{}

func init() {
	f := BKaffa咖法.(*咖法KaffaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaffa",
		TitleName: "咖法",
		TitleCode: "b_kaffa",
	}
}
