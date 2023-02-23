package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮恩扎PienzaBarony struct {
	feud.BaseBarony
}

var BPienza皮恩扎 feud.Barony = &皮恩扎PienzaBarony{}

func init() {
	f := BPienza皮恩扎.(*皮恩扎PienzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pienza",
		TitleName: "皮恩扎",
		TitleCode: "b_pienza",
	}
}
