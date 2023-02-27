package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉克卡KapalakBarony struct {
	feud.BaseBarony
}

var BKapalak卡拉克卡 feud.Barony = &卡拉克卡KapalakBarony{}

func init() {
    f := BKapalak卡拉克卡.(*卡拉克卡KapalakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapalak",
		TitleName: "卡拉克卡",
		TitleCode: "b_kapalak",
	}
}
