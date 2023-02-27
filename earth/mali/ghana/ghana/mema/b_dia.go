package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚DiaBarony struct {
	feud.BaseBarony
}

var BDia迪亚 feud.Barony = &迪亚DiaBarony{}

func init() {
    f := BDia迪亚.(*迪亚DiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dia",
		TitleName: "迪亚",
		TitleCode: "b_dia",
	}
}
