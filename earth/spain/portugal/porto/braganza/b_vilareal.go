package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷阿尔城VilarealBarony struct {
	feud.BaseBarony
}

var BVilareal雷阿尔城 feud.Barony = &雷阿尔城VilarealBarony{}

func init() {
    f := BVilareal雷阿尔城.(*雷阿尔城VilarealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vilareal",
		TitleName: "雷阿尔城",
		TitleCode: "b_vilareal",
	}
}
