package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔斯兰GilslandBarony struct {
	feud.BaseBarony
}

var BGilsland吉尔斯兰 feud.Barony = &吉尔斯兰GilslandBarony{}

func init() {
    f := BGilsland吉尔斯兰.(*吉尔斯兰GilslandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gilsland",
		TitleName: "吉尔斯兰",
		TitleCode: "b_gilsland",
	}
}
