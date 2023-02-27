package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗韶梨RajauliBarony struct {
	feud.BaseBarony
}

var BRajauli罗韶梨 feud.Barony = &罗韶梨RajauliBarony{}

func init() {
    f := BRajauli罗韶梨.(*罗韶梨RajauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajauli",
		TitleName: "罗韶梨",
		TitleCode: "b_rajauli",
	}
}
