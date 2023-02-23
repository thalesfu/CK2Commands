package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈索隆吉MesolongiBarony struct {
	feud.BaseBarony
}

var BMesolongi迈索隆吉 feud.Barony = &迈索隆吉MesolongiBarony{}

func init() {
	f := BMesolongi迈索隆吉.(*迈索隆吉MesolongiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mesolongi",
		TitleName: "迈索隆吉",
		TitleCode: "b_mesolongi",
	}
}
