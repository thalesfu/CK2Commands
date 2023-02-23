package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍比凯ChebikaBarony struct {
	feud.BaseBarony
}

var BChebika舍比凯 feud.Barony = &舍比凯ChebikaBarony{}

func init() {
	f := BChebika舍比凯.(*舍比凯ChebikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chebika",
		TitleName: "舍比凯",
		TitleCode: "b_chebika",
	}
}
