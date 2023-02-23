package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达利阿DhalaBarony struct {
	feud.BaseBarony
}

var BDhala达利阿 feud.Barony = &达利阿DhalaBarony{}

func init() {
	f := BDhala达利阿.(*达利阿DhalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhala",
		TitleName: "达利阿",
		TitleCode: "b_dhala",
	}
}
