package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫斯比HesbyBarony struct {
	feud.BaseBarony
}

var BHesby赫斯比 feud.Barony = &赫斯比HesbyBarony{}

func init() {
    f := BHesby赫斯比.(*赫斯比HesbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hesby",
		TitleName: "赫斯比",
		TitleCode: "b_hesby",
	}
}
