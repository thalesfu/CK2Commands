package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨卡MasakaBarony struct {
	feud.BaseBarony
}

var BMasaka马萨卡 feud.Barony = &马萨卡MasakaBarony{}

func init() {
	f := BMasaka马萨卡.(*马萨卡MasakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masaka",
		TitleName: "马萨卡",
		TitleCode: "b_masaka",
	}
}
