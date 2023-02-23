package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔KilBarony struct {
	feud.BaseBarony
}

var BKil希尔 feud.Barony = &希尔KilBarony{}

func init() {
	f := BKil希尔.(*希尔KilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kil",
		TitleName: "希尔",
		TitleCode: "b_kil",
	}
}
