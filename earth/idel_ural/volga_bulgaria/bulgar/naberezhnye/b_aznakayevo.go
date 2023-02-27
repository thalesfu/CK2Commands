package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿兹纳卡耶沃AznakayevoBarony struct {
	feud.BaseBarony
}

var BAznakayevo阿兹纳卡耶沃 feud.Barony = &阿兹纳卡耶沃AznakayevoBarony{}

func init() {
    f := BAznakayevo阿兹纳卡耶沃.(*阿兹纳卡耶沃AznakayevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aznakayevo",
		TitleName: "阿兹纳卡耶沃",
		TitleCode: "b_aznakayevo",
	}
}
