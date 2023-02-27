package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷瑟哈ReeshahBarony struct {
	feud.BaseBarony
}

var BReeshah雷瑟哈 feud.Barony = &雷瑟哈ReeshahBarony{}

func init() {
    f := BReeshah雷瑟哈.(*雷瑟哈ReeshahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reeshah",
		TitleName: "雷瑟哈",
		TitleCode: "b_reeshah",
	}
}
