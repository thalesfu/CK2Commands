package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀罗DhrolBarony struct {
	feud.BaseBarony
}

var BDhrol陀罗 feud.Barony = &陀罗DhrolBarony{}

func init() {
	f := BDhrol陀罗.(*陀罗DhrolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhrol",
		TitleName: "陀罗",
		TitleCode: "b_dhrol",
	}
}
