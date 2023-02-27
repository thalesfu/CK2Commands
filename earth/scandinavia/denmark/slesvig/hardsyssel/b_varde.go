package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦德VardeBarony struct {
	feud.BaseBarony
}

var BVarde瓦德 feud.Barony = &瓦德VardeBarony{}

func init() {
    f := BVarde瓦德.(*瓦德VardeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varde",
		TitleName: "瓦德",
		TitleCode: "b_varde",
	}
}
