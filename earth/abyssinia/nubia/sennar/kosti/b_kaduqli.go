package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡杜格利KaduqliBarony struct {
	feud.BaseBarony
}

var BKaduqli卡杜格利 feud.Barony = &卡杜格利KaduqliBarony{}

func init() {
	f := BKaduqli卡杜格利.(*卡杜格利KaduqliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaduqli",
		TitleName: "卡杜格利",
		TitleCode: "b_kaduqli",
	}
}
