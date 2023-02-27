package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 囊伽梨NangaliBarony struct {
	feud.BaseBarony
}

var BNangali囊伽梨 feud.Barony = &囊伽梨NangaliBarony{}

func init() {
    f := BNangali囊伽梨.(*囊伽梨NangaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nangali",
		TitleName: "囊伽梨",
		TitleCode: "b_nangali",
	}
}
