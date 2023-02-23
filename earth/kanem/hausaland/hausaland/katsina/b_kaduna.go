package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡杜纳KadunaBarony struct {
	feud.BaseBarony
}

var BKaduna卡杜纳 feud.Barony = &卡杜纳KadunaBarony{}

func init() {
	f := BKaduna卡杜纳.(*卡杜纳KadunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaduna",
		TitleName: "卡杜纳",
		TitleCode: "b_kaduna",
	}
}
