package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫格斯比HogsbyBarony struct {
	feud.BaseBarony
}

var BHogsby赫格斯比 feud.Barony = &赫格斯比HogsbyBarony{}

func init() {
	f := BHogsby赫格斯比.(*赫格斯比HogsbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hogsby",
		TitleName: "赫格斯比",
		TitleCode: "b_hogsby",
	}
}
