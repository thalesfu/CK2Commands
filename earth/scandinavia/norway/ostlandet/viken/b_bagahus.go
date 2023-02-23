package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格赫斯BagahusBarony struct {
	feud.BaseBarony
}

var BBagahus巴格赫斯 feud.Barony = &巴格赫斯BagahusBarony{}

func init() {
	f := BBagahus巴格赫斯.(*巴格赫斯BagahusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagahus",
		TitleName: "巴格赫斯",
		TitleCode: "b_bagahus",
	}
}
