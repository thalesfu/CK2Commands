package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷斯伦LathrunBarony struct {
	feud.BaseBarony
}

var BLathrun雷斯伦 feud.Barony = &雷斯伦LathrunBarony{}

func init() {
	f := BLathrun雷斯伦.(*雷斯伦LathrunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lathrun",
		TitleName: "雷斯伦",
		TitleCode: "b_lathrun",
	}
}
