package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科什基KoshkiBarony struct {
	feud.BaseBarony
}

var BKoshki科什基 feud.Barony = &科什基KoshkiBarony{}

func init() {
    f := BKoshki科什基.(*科什基KoshkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koshki",
		TitleName: "科什基",
		TitleCode: "b_koshki",
	}
}
