package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴卡拉BakaraBarony struct {
	feud.BaseBarony
}

var BBakara巴卡拉 feud.Barony = &巴卡拉BakaraBarony{}

func init() {
    f := BBakara巴卡拉.(*巴卡拉BakaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakara",
		TitleName: "巴卡拉",
		TitleCode: "b_bakara",
	}
}
