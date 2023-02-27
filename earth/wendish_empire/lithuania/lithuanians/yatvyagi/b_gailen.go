package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖伦GailenBarony struct {
	feud.BaseBarony
}

var BGailen盖伦 feud.Barony = &盖伦GailenBarony{}

func init() {
    f := BGailen盖伦.(*盖伦GailenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gailen",
		TitleName: "盖伦",
		TitleCode: "b_gailen",
	}
}
