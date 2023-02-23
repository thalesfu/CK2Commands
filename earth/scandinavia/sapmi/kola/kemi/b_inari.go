package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊纳里InariBarony struct {
	feud.BaseBarony
}

var BInari伊纳里 feud.Barony = &伊纳里InariBarony{}

func init() {
	f := BInari伊纳里.(*伊纳里InariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inari",
		TitleName: "伊纳里",
		TitleCode: "b_inari",
	}
}
