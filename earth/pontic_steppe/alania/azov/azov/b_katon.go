package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡通KatonBarony struct {
	feud.BaseBarony
}

var BKaton卡通 feud.Barony = &卡通KatonBarony{}

func init() {
    f := BKaton卡通.(*卡通KatonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katon",
		TitleName: "卡通",
		TitleCode: "b_katon",
	}
}
