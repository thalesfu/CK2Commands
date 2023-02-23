package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哥本哈根KobenhavnBarony struct {
	feud.BaseBarony
}

var BKobenhavn哥本哈根 feud.Barony = &哥本哈根KobenhavnBarony{}

func init() {
	f := BKobenhavn哥本哈根.(*哥本哈根KobenhavnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kobenhavn",
		TitleName: "哥本哈根",
		TitleCode: "b_kobenhavn",
	}
}
