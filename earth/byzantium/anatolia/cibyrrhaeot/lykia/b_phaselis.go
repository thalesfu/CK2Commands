package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法塞利斯PhaselisBarony struct {
	feud.BaseBarony
}

var BPhaselis法塞利斯 feud.Barony = &法塞利斯PhaselisBarony{}

func init() {
    f := BPhaselis法塞利斯.(*法塞利斯PhaselisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phaselis",
		TitleName: "法塞利斯",
		TitleCode: "b_phaselis",
	}
}
