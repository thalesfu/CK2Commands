package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达哈尔吉赫纳Dahar_jihenaBarony struct {
	feud.BaseBarony
}

var BDahar_jihena达哈尔吉赫纳 feud.Barony = &达哈尔吉赫纳Dahar_jihenaBarony{}

func init() {
    f := BDahar_jihena达哈尔吉赫纳.(*达哈尔吉赫纳Dahar_jihenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dahar_jihena",
		TitleName: "达哈尔吉赫纳",
		TitleCode: "b_dahar_jihena",
	}
}
