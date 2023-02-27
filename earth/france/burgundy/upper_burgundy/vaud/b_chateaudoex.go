package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代堡ChateaudoexBarony struct {
	feud.BaseBarony
}

var BChateaudoex代堡 feud.Barony = &代堡ChateaudoexBarony{}

func init() {
    f := BChateaudoex代堡.(*代堡ChateaudoexBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateaudoex",
		TitleName: "代堡",
		TitleCode: "b_chateaudoex",
	}
}
