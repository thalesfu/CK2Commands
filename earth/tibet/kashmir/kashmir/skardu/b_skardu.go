package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡卡都SkarduBarony struct {
	feud.BaseBarony
}

var BSkardu锡卡都 feud.Barony = &锡卡都SkarduBarony{}

func init() {
    f := BSkardu锡卡都.(*锡卡都SkarduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skardu",
		TitleName: "锡卡都",
		TitleCode: "b_skardu",
	}
}
