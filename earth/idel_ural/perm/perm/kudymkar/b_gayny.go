package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖内GaynyBarony struct {
	feud.BaseBarony
}

var BGayny盖内 feud.Barony = &盖内GaynyBarony{}

func init() {
    f := BGayny盖内.(*盖内GaynyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gayny",
		TitleName: "盖内",
		TitleCode: "b_gayny",
	}
}
