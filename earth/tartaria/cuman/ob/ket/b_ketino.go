package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克蒂诺KetinoBarony struct {
	feud.BaseBarony
}

var BKetino克蒂诺 feud.Barony = &克蒂诺KetinoBarony{}

func init() {
	f := BKetino克蒂诺.(*克蒂诺KetinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ketino",
		TitleName: "克蒂诺",
		TitleCode: "b_ketino",
	}
}
