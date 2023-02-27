package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶尔西卡JersikaBarony struct {
	feud.BaseBarony
}

var BJersika耶尔西卡 feud.Barony = &耶尔西卡JersikaBarony{}

func init() {
    f := BJersika耶尔西卡.(*耶尔西卡JersikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jersika",
		TitleName: "耶尔西卡",
		TitleCode: "b_jersika",
	}
}
