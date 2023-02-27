package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法努瓦迪格FanoidigBarony struct {
	feud.BaseBarony
}

var BFanoidig法努瓦迪格 feud.Barony = &法努瓦迪格FanoidigBarony{}

func init() {
    f := BFanoidig法努瓦迪格.(*法努瓦迪格FanoidigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fanoidig",
		TitleName: "法努瓦迪格",
		TitleCode: "b_fanoidig",
	}
}
