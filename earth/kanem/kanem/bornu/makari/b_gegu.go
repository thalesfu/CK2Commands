package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖古GeguBarony struct {
	feud.BaseBarony
}

var BGegu盖古 feud.Barony = &盖古GeguBarony{}

func init() {
	f := BGegu盖古.(*盖古GeguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gegu",
		TitleName: "盖古",
		TitleCode: "b_gegu",
	}
}
