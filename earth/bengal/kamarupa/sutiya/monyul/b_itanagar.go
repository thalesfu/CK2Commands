package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊塔那噶ItanagarBarony struct {
	feud.BaseBarony
}

var BItanagar伊塔那噶 feud.Barony = &伊塔那噶ItanagarBarony{}

func init() {
	f := BItanagar伊塔那噶.(*伊塔那噶ItanagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "itanagar",
		TitleName: "伊塔那噶",
		TitleCode: "b_itanagar",
	}
}
