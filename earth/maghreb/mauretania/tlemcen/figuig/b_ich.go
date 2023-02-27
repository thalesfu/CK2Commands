package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊什IchBarony struct {
	feud.BaseBarony
}

var BIch伊什 feud.Barony = &伊什IchBarony{}

func init() {
    f := BIch伊什.(*伊什IchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ich",
		TitleName: "伊什",
		TitleCode: "b_ich",
	}
}
