package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰恰姆加ChachamgaBarony struct {
	feud.BaseBarony
}

var BChachamga恰恰姆加 feud.Barony = &恰恰姆加ChachamgaBarony{}

func init() {
	f := BChachamga恰恰姆加.(*恰恰姆加ChachamgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chachamga",
		TitleName: "恰恰姆加",
		TitleCode: "b_chachamga",
	}
}
