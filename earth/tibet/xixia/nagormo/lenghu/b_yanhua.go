package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盐化YanhuaBarony struct {
	feud.BaseBarony
}

var BYanhua盐化 feud.Barony = &盐化YanhuaBarony{}

func init() {
	f := BYanhua盐化.(*盐化YanhuaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yanhua",
		TitleName: "盐化",
		TitleCode: "b_yanhua",
	}
}
