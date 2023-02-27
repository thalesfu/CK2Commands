package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐齐格穆林Tsetseg_moriinBarony struct {
	feud.BaseBarony
}

var BTsetseg_moriin齐齐格穆林 feud.Barony = &齐齐格穆林Tsetseg_moriinBarony{}

func init() {
    f := BTsetseg_moriin齐齐格穆林.(*齐齐格穆林Tsetseg_moriinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsetseg_moriin",
		TitleName: "齐齐格穆林",
		TitleCode: "b_tsetseg_moriin",
	}
}
