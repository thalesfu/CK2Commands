package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦泰希LentekhiBarony struct {
	feud.BaseBarony
}

var BLentekhi伦泰希 feud.Barony = &伦泰希LentekhiBarony{}

func init() {
	f := BLentekhi伦泰希.(*伦泰希LentekhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lentekhi",
		TitleName: "伦泰希",
		TitleCode: "b_lentekhi",
	}
}
