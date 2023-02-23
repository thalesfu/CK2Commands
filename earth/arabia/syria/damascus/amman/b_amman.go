package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安曼AmmanBarony struct {
	feud.BaseBarony
}

var BAmman安曼 feud.Barony = &安曼AmmanBarony{}

func init() {
	f := BAmman安曼.(*安曼AmmanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amman",
		TitleName: "安曼",
		TitleCode: "b_amman",
	}
}
