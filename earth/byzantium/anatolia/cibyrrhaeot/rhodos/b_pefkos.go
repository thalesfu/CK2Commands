package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩夫科斯PefkosBarony struct {
	feud.BaseBarony
}

var BPefkos佩夫科斯 feud.Barony = &佩夫科斯PefkosBarony{}

func init() {
	f := BPefkos佩夫科斯.(*佩夫科斯PefkosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pefkos",
		TitleName: "佩夫科斯",
		TitleCode: "b_pefkos",
	}
}
