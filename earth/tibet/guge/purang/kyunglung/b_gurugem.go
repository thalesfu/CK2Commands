package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古入江GurugemBarony struct {
	feud.BaseBarony
}

var BGurugem古入江 feud.Barony = &古入江GurugemBarony{}

func init() {
    f := BGurugem古入江.(*古入江GurugemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurugem",
		TitleName: "古入江",
		TitleCode: "b_gurugem",
	}
}
