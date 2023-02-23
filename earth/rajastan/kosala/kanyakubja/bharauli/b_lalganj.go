package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗梨犍阇LalganjBarony struct {
	feud.BaseBarony
}

var BLalganj罗梨犍阇 feud.Barony = &罗梨犍阇LalganjBarony{}

func init() {
	f := BLalganj罗梨犍阇.(*罗梨犍阇LalganjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lalganj",
		TitleName: "罗梨犍阇",
		TitleCode: "b_lalganj",
	}
}
