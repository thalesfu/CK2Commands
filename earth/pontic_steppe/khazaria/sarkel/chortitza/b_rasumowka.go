package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉祖莫夫卡RasumowkaBarony struct {
	feud.BaseBarony
}

var BRasumowka拉祖莫夫卡 feud.Barony = &拉祖莫夫卡RasumowkaBarony{}

func init() {
    f := BRasumowka拉祖莫夫卡.(*拉祖莫夫卡RasumowkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rasumowka",
		TitleName: "拉祖莫夫卡",
		TitleCode: "b_rasumowka",
	}
}
