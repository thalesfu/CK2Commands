package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马西纳MacinaBarony struct {
	feud.BaseBarony
}

var BMacina马西纳 feud.Barony = &马西纳MacinaBarony{}

func init() {
	f := BMacina马西纳.(*马西纳MacinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "macina",
		TitleName: "马西纳",
		TitleCode: "b_macina",
	}
}
