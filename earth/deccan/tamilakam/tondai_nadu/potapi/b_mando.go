package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼度MandoBarony struct {
	feud.BaseBarony
}

var BMando曼度 feud.Barony = &曼度MandoBarony{}

func init() {
    f := BMando曼度.(*曼度MandoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mando",
		TitleName: "曼度",
		TitleCode: "b_mando",
	}
}
