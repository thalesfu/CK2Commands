package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾科维察DjakovicaBarony struct {
	feud.BaseBarony
}

var BDjakovica贾科维察 feud.Barony = &贾科维察DjakovicaBarony{}

func init() {
    f := BDjakovica贾科维察.(*贾科维察DjakovicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djakovica",
		TitleName: "贾科维察",
		TitleCode: "b_djakovica",
	}
}
