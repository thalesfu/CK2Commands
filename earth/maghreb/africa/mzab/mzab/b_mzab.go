package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆扎卜MzabBarony struct {
	feud.BaseBarony
}

var BMzab姆扎卜 feud.Barony = &姆扎卜MzabBarony{}

func init() {
	f := BMzab姆扎卜.(*姆扎卜MzabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mzab",
		TitleName: "姆扎卜",
		TitleCode: "b_mzab",
	}
}
