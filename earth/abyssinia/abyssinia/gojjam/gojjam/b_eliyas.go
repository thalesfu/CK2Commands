package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊利亚斯EliyasBarony struct {
	feud.BaseBarony
}

var BEliyas伊利亚斯 feud.Barony = &伊利亚斯EliyasBarony{}

func init() {
	f := BEliyas伊利亚斯.(*伊利亚斯EliyasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eliyas",
		TitleName: "伊利亚斯",
		TitleCode: "b_eliyas",
	}
}
