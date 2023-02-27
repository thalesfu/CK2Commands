package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩伽RamagramaBarony struct {
	feud.BaseBarony
}

var BRamagrama罗摩伽 feud.Barony = &罗摩伽RamagramaBarony{}

func init() {
    f := BRamagrama罗摩伽.(*罗摩伽RamagramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramagrama",
		TitleName: "罗摩伽",
		TitleCode: "b_ramagrama",
	}
}
