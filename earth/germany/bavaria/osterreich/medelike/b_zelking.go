package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 策尔京ZelkingBarony struct {
	feud.BaseBarony
}

var BZelking策尔京 feud.Barony = &策尔京ZelkingBarony{}

func init() {
	f := BZelking策尔京.(*策尔京ZelkingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zelking",
		TitleName: "策尔京",
		TitleCode: "b_zelking",
	}
}
