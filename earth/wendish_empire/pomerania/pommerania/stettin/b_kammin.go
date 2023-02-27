package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡明KamminBarony struct {
	feud.BaseBarony
}

var BKammin卡明 feud.Barony = &卡明KamminBarony{}

func init() {
    f := BKammin卡明.(*卡明KamminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kammin",
		TitleName: "卡明",
		TitleCode: "b_kammin",
	}
}
