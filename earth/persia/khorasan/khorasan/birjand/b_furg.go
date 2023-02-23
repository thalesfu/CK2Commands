package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富尔格FurgBarony struct {
	feud.BaseBarony
}

var BFurg富尔格 feud.Barony = &富尔格FurgBarony{}

func init() {
	f := BFurg富尔格.(*富尔格FurgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furg",
		TitleName: "富尔格",
		TitleCode: "b_furg",
	}
}
