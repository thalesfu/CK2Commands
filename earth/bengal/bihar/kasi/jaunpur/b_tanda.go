package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 檀陀TandaBarony struct {
	feud.BaseBarony
}

var BTanda檀陀 feud.Barony = &檀陀TandaBarony{}

func init() {
	f := BTanda檀陀.(*檀陀TandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanda",
		TitleName: "檀陀",
		TitleCode: "b_tanda",
	}
}
